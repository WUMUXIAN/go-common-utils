// Package srp implements the Secure Remote Password following SRP-6a protocol design.
package srp

// This implementation follows the Secure Remote Password Design, latest SRP-6a.

// SRP Protocol Design

// SRP is the newest addition to a new class of strong authentication protocols that resist all the well-known passive and active attacks over the network.
// SRP borrows some elements from other key-exchange and identification protcols and adds some subtle modifications and refinements.
// The result is a protocol that preserves the strength and efficiency of the EKE family protocols while fixing some of their shortcomings.
// The following is a description of SRP-6 and 6a, the latest versions of SRP:

//   N    A large safe prime (N = 2q+1, where q is prime)
//        All arithmetic is done modulo N.
//   g    A generator modulo N
//   k    Multiplier parameter (k = H(N, g) in SRP-6a, k = 3 for legacy SRP-6)
//   s    User's salt
//   I    Username
//   p    Cleartext Password
//   H()  One-way hash function
//   ^    (Modular) Exponentiation
//   u    Random scrambling parameter
//   a,b  Secret ephemeral values
//   A,B  Public ephemeral values
//   x    Private key (derived from p and s)
//   v    Password verifier
// The host stores passwords using the following formula:
//   x = H(s, p)               (s is chosen randomly)
//   v = g^x                   (computes password verifier)
// The host then keeps {I, s, v} in its password database. The authentication protocol itself goes as follows:
// User -> Host:  I, A = g^a                  (identifies self, a = random number)
// Host -> User:  s, B = kv + g^b             (sends salt, b = random number)

//         Both:  u = H(A, B)

//         User:  x = H(s, p)                 (user enters password)
//         User:  S = (B - kg^x) ^ (a + ux)   (computes session key)
//         User:  K = H(S)

//         Host:  S = (Av^u) ^ b              (computes session key)
//         Host:  K = H(S)
// Now the two parties have a shared, strong session key K. To complete authentication, they need to prove to each other that their keys match. One possible way:
// User -> Host:  M = H(H(N) xor H(g), H(I), s, A, B, K)
// Host -> User:  H(A, M, K)
// The two parties also employ the following safeguards:
// The user will abort if he receives B == 0 (mod N) or u == 0.
// The host will abort if it detects that A == 0 (mod N).
// The user must show his proof of K first. If the server detects that the user's proof is incorrect, it must abort without showing its own proof of K.

// In my implementation:

// H = SHA256()
// k = H(N, g)
// I = User's email address
// i = lower case of user's email address
// p = NFKD(trim(p)) -- The password is trimmed and them normalized before passed to calculate x
// s = HKDF(s, i, I, 32)
// x = PBKDF2_HAMC_SHA256(p, s, 100000)
// M (client -> server) = HAMC-SHA256(K, A, B, I, s, N, g)
// Proof (server -> client) = HMAC-SHA256(K, M)

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/WUMUXIAN/go-common-utils/codec"
	"github.com/WUMUXIAN/go-common-utils/cryptowrapper"
	"golang.org/x/text/unicode/norm"
)

var gNStrMap = map[int][2]string{
	1024: {"2", "0xEEAF0AB9ADB38DD69C33F80AFA8FC5E86072618775FF3C0B9EA2314C9C256576D674DF7496EA81D3383B4813D692C6E0E0D5D8E250B98BE48E495C1D6089DAD15DC7D7B46154D6B6CE8EF4AD69B15D4982559B297BCF1885C529F566660E57EC68EDBC3C05726CC02FD4CBF4976EAA9AFD5138FE8376435B9FC61D2FC0EB06E3"},
	1536: {"2", "0x9DEF3CAFB939277AB1F12A8617A47BBBDBA51DF499AC4C80BEEEA9614B19CC4D5F4F5F556E27CBDE51C6A94BE4607A291558903BA0D0F84380B655BB9A22E8DCDF028A7CEC67F0D08134B1C8B97989149B609E0BE3BAB63D47548381DBC5B1FC764E3F4B53DD9DA1158BFD3E2B9C8CF56EDF019539349627DB2FD53D24B7C48665772E437D6C7F8CE442734AF7CCB7AE837C264AE3A9BEB87F8A2FE9B8B5292E5A021FFF5E91479E8CE7A28C2442C6F315180F93499A234DCF76E3FED135F9BB"},
	2048: {"2", "0xAC6BDB41324A9A9BF166DE5E1389582FAF72B6651987EE07FC3192943DB56050A37329CBB4A099ED8193E0757767A13DD52312AB4B03310DCD7F48A9DA04FD50E8083969EDB767B0CF6095179A163AB3661A05FBD5FAAAE82918A9962F0B93B855F97993EC975EEAA80D740ADBF4FF747359D041D5C33EA71D281E446B14773BCA97B43A23FB801676BD207A436C6481F1D2B9078717461A5B9D32E688F87748544523B524B0D57D5EA77A2775D2ECFA032CFBDBF52FB3786160279004E57AE6AF874E7303CE53299CCC041C7BC308D82A5698F3A8D0C38271AE35F8E9DBFBB694B5C803D89F7AE435DE236D525F54759B65E372FCD68EF20FA7111F9E4AFF73"},
	3072: {"2", "0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AAAC42DAD33170D04507A33A85521ABDF1CBA64ECFB850458DBEF0A8AEA71575D060C7DB3970F85A6E1E4C7ABF5AE8CDB0933D71E8C94E04A25619DCEE3D2261AD2EE6BF12FFA06D98A0864D87602733EC86A64521F2B18177B200CBBE117577A615D6C770988C0BAD946E208E24FA074E5AB3143DB5BFCE0FD108E4B82D120A93AD2CAFFFFFFFFFFFFFFFF"},
	4096: {"5", "0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AAAC42DAD33170D04507A33A85521ABDF1CBA64ECFB850458DBEF0A8AEA71575D060C7DB3970F85A6E1E4C7ABF5AE8CDB0933D71E8C94E04A25619DCEE3D2261AD2EE6BF12FFA06D98A0864D87602733EC86A64521F2B18177B200CBBE117577A615D6C770988C0BAD946E208E24FA074E5AB3143DB5BFCE0FD108E4B82D120A92108011A723C12A787E6D788719A10BDBA5B2699C327186AF4E23C1A946834B6150BDA2583E9CA2AD44CE8DBBBC2DB04DE8EF92E8EFC141FBECAA6287C59474E6BC05D99B2964FA090C3A2233BA186515BE7ED1F612970CEE2D7AFB81BDD762170481CD0069127D5B05AA993B4EA988D8FDDC186FFB7DC90A6C08F4DF435C934063199FFFFFFFFFFFFFFFF"},
	6144: {"5", "0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AAAC42DAD33170D04507A33A85521ABDF1CBA64ECFB850458DBEF0A8AEA71575D060C7DB3970F85A6E1E4C7ABF5AE8CDB0933D71E8C94E04A25619DCEE3D2261AD2EE6BF12FFA06D98A0864D87602733EC86A64521F2B18177B200CBBE117577A615D6C770988C0BAD946E208E24FA074E5AB3143DB5BFCE0FD108E4B82D120A92108011A723C12A787E6D788719A10BDBA5B2699C327186AF4E23C1A946834B6150BDA2583E9CA2AD44CE8DBBBC2DB04DE8EF92E8EFC141FBECAA6287C59474E6BC05D99B2964FA090C3A2233BA186515BE7ED1F612970CEE2D7AFB81BDD762170481CD0069127D5B05AA993B4EA988D8FDDC186FFB7DC90A6C08F4DF435C93402849236C3FAB4D27C7026C1D4DCB2602646DEC9751E763DBA37BDF8FF9406AD9E530EE5DB382F413001AEB06A53ED9027D831179727B0865A8918DA3EDBEBCF9B14ED44CE6CBACED4BB1BDB7F1447E6CC254B332051512BD7AF426FB8F401378CD2BF5983CA01C64B92ECF032EA15D1721D03F482D7CE6E74FEF6D55E702F46980C82B5A84031900B1C9E59E7C97FBEC7E8F323A97A7E36CC88BE0F1D45B7FF585AC54BD407B22B4154AACC8F6D7EBF48E1D814CC5ED20F8037E0A79715EEF29BE32806A1D58BB7C5DA76F550AA3D8A1FBFF0EB19CCB1A313D55CDA56C9EC2EF29632387FE8D76E3C0468043E8F663F4860EE12BF2D5B0B7474D6E694F91E6DCC4024FFFFFFFFFFFFFFFF"},
	8192: {"5", "0xFFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AAAC42DAD33170D04507A33A85521ABDF1CBA64ECFB850458DBEF0A8AEA71575D060C7DB3970F85A6E1E4C7ABF5AE8CDB0933D71E8C94E04A25619DCEE3D2261AD2EE6BF12FFA06D98A0864D87602733EC86A64521F2B18177B200CBBE117577A615D6C770988C0BAD946E208E24FA074E5AB3143DB5BFCE0FD108E4B82D120A92108011A723C12A787E6D788719A10BDBA5B2699C327186AF4E23C1A946834B6150BDA2583E9CA2AD44CE8DBBBC2DB04DE8EF92E8EFC141FBECAA6287C59474E6BC05D99B2964FA090C3A2233BA186515BE7ED1F612970CEE2D7AFB81BDD762170481CD0069127D5B05AA993B4EA988D8FDDC186FFB7DC90A6C08F4DF435C93402849236C3FAB4D27C7026C1D4DCB2602646DEC9751E763DBA37BDF8FF9406AD9E530EE5DB382F413001AEB06A53ED9027D831179727B0865A8918DA3EDBEBCF9B14ED44CE6CBACED4BB1BDB7F1447E6CC254B332051512BD7AF426FB8F401378CD2BF5983CA01C64B92ECF032EA15D1721D03F482D7CE6E74FEF6D55E702F46980C82B5A84031900B1C9E59E7C97FBEC7E8F323A97A7E36CC88BE0F1D45B7FF585AC54BD407B22B4154AACC8F6D7EBF48E1D814CC5ED20F8037E0A79715EEF29BE32806A1D58BB7C5DA76F550AA3D8A1FBFF0EB19CCB1A313D55CDA56C9EC2EF29632387FE8D76E3C0468043E8F663F4860EE12BF2D5B0B7474D6E694F91E6DBE115974A3926F12FEE5E438777CB6A932DF8CD8BEC4D073B931BA3BC832B68D9DD300741FA7BF8AFC47ED2576F6936BA424663AAB639C5AE4F5683423B4742BF1C978238F16CBE39D652DE3FDB8BEFC848AD922222E04A4037C0713EB57A81A23F0C73473FC646CEA306B4BCBC8862F8385DDFA9D4B7FA2C087E879683303ED5BDD3A062B3CF5B3A278A66D2A13F83F44F82DDF310EE074AB6A364597E899A0255DC164F31CC50846851DF9AB48195DED7EA1B1D510BD7EE74D73FAF36BC31ECFA268359046F4EB879F924009438B481C6CD7889A002ED5EE382BC9190DA6FC026E479558E4475677E9AA9E3050E2765694DFC81F56E880B96E7160C980DD98EDD3DFFFFFFFFFFFFFFFFF"},
}

var hashFunc = sha256.New

type prime struct {
	g *big.Int
	N *big.Int
}

var gNMap map[int]prime

func init() {
	gNMap = make(map[int]prime)

	for bits, arr := range gNStrMap {
		g, _ := big.NewInt(0).SetString(arr[0], 10)
		n, _ := big.NewInt(0).SetString(arr[1], 0)
		gNMap[bits] = prime{g: g, N: n}
	}
}

// SRPClient defines a SRP client
type SRPClient struct {
	prime
	p []byte   // password
	E []byte   // email
	S []byte   // salt
	a *big.Int // random
	A *big.Int // Public version of a
	k *big.Int // k = H(N, g)
	x *big.Int

	K []byte // shared session key
	M []byte // The verification bytes used for prove client
}

// NewSRPClient creates a SRP client
func NewSRPClient(email, password string, primeBit int) (c *SRPClient, err error) {
	c = new(SRPClient)
	err = c.init(email, password, primeBit)
	return
}

func (c *SRPClient) init(email, password string, primeBit int) (err error) {
	prime, ok := gNMap[primeBit]
	if !ok || 0 == big.NewInt(0).Cmp(prime.N) {
		return errors.New("invalid bits")
	}

	c.g = prime.g
	c.N = prime.N
	c.p = norm.NFKD.Bytes([]byte(strings.TrimSpace(password)))
	c.E = []byte(email)
	c.a = cryptowrapper.RandBigInt(primeBit)
	c.A = big.NewInt(0).Exp(prime.g, c.a, prime.N)
	c.k = codec.GetHash(codec.SHA256, prime.N.Bytes(), prime.g.Bytes()).BigInt()

	salt := cryptowrapper.RandBytes(32)
	salt, err = cryptowrapper.DeriveKey(hashFunc, salt, []byte(email), []byte(strings.ToLower(email)), 32)

	// fmt.Println("g: ", c.g.Bytes())
	// fmt.Println("N: ", codec.ToByteArray(c.N.Bytes()).Base64())
	// fmt.Println("k length: ", len(c.k.Bytes())*8)
	// fmt.Println("k: ", c.k.Bytes())
	// fmt.Println("a length: ", len(c.a.Bytes())*8)
	// fmt.Println("a: ", codec.ToByteArray(c.a.Bytes()).Base64())
	// fmt.Println("A length: ", len(c.A.Bytes())*8)
	// fmt.Println("A: ", codec.ToByteArray(c.A.Bytes()).Base64())
	// fmt.Println("Random Salt: ", codec.ToByteArray(salt).Base64())
	// fmt.Println("SRP-x Salt: ", codec.ToByteArray(salt).Base64())

	c.S = salt
	return
}

// GetX calculates the SRP X
func (c *SRPClient) GetX() {
	c.x = codec.ToByteArray(cryptowrapper.PBKDF2HMACSHA256(c.p, c.S, 100000, 32)).BigInt()
}

// GetVerifier calculates the SRP Verifier
func (c *SRPClient) GetVerifier() (v *big.Int) {
	// v = g^x
	v = big.NewInt(0).Exp(c.g, c.x, c.N)
	return
}

// Calculate performs the SRP calculation
func (c *SRPClient) Calculate(B *big.Int, salt []byte) (err error) {
	if big.NewInt(0).Cmp(big.NewInt(0).Mod(B, c.N)) == 0 {
		err = fmt.Errorf("invalid server public key B=%x", B)
		return
	}

	u := codec.GetHash(codec.SHA256, c.A.Bytes(), B.Bytes()).BigInt()

	tmp1 := big.NewInt(0).Mul(c.k, big.NewInt(0).Exp(c.g, c.x, c.N))
	tmp2 := big.NewInt(0).Sub(B, tmp1)
	tmp3 := big.NewInt(0).Add(c.a, big.NewInt(0).Mul(u, c.x))
	secret := big.NewInt(0).Exp(tmp2, tmp3, c.N)

	// fmt.Println("u = ", len(u.Bytes()) * 8)
	// fmt.Println("k*(g^x mod N) = ", len(tmp1.Bytes()) * 8)
	// fmt.Println("B - k*(g^x mod N) = ", len(tmp2.Bytes()) * 8)
	// fmt.Println("a + u*x = ", len(tmp3.Bytes()) * 8)
	// fmt.Println("(B - k*(g^x mod N)) ^ (a + u*x) = ", len(secret.Bytes()) * 8)

	c.K = codec.GetHash(codec.SHA256, secret.Bytes())
	c.M = cryptowrapper.HMACSHA256(c.K, c.A.Bytes(), B.Bytes(), c.E, salt, c.N.Bytes(), c.g.Bytes())

	return
}

// CheckServer checks whether the proof given by server is correct.
func (c *SRPClient) CheckServer(proof []byte) (err error) {
	h := cryptowrapper.HMACSHA256(c.K, c.M)
	if bytes.Equal(h, proof) {
		return nil
	}
	return errors.New("server authentication failed")
}

// SRPServer defines the SRP server
type SRPServer struct {
	prime
	S []byte   // salt
	e []byte   // email
	v *big.Int // verifier
	b *big.Int // random
	B *big.Int // Public version of b

	K []byte // Commonly shared session key.
	M []byte // Key to verify client
}

// NewSRPServer creates a new SRP server
func NewSRPServer(email, salt []byte, verifier *big.Int, A *big.Int, primeBit int) (s *SRPServer, err error) {
	s = new(SRPServer)
	err = s.init(email, salt, verifier, A, primeBit)
	return
}

func (s *SRPServer) init(email, salt []byte, verifier *big.Int, A *big.Int, primeBit int) (err error) {
	prime, ok := gNMap[primeBit]
	if !ok || 0 == big.NewInt(0).Cmp(prime.N) {
		return errors.New("invalid bits")
	}

	s.v = verifier
	s.g = prime.g
	s.N = prime.N
	s.e = email
	s.S = salt

	s.b = cryptowrapper.RandBigInt(primeBit)
	k := codec.GetHash(codec.SHA256, prime.N.Bytes(), prime.g.Bytes()).BigInt()
	s.B = big.NewInt(0).Add(big.NewInt(0).Mul(k, s.v), big.NewInt(0).Exp(s.g, s.b, s.N))

	// fmt.Println("big.NewInt(0).Mul(k, s.v) ", codec.ToByteArray(big.NewInt(0).Mul(k, s.v).Bytes()).Base64())
	// fmt.Println("big.NewInt(0).Exp(s.g, s.b, s.N) ", codec.ToByteArray(big.NewInt(0).Exp(s.g, s.b, s.N).Bytes()).Base64())
	//
	// fmt.Println("b length: ", len(s.b.Bytes())*8)
	// fmt.Println("b: ", codec.ToByteArray(s.b.Bytes()).Base64())
	// fmt.Println("B length: ", len(s.B.Bytes())*8)
	// fmt.Println("B: ", codec.ToByteArray(s.B.Bytes()).Base64())

	u := codec.GetHash(codec.SHA256, A.Bytes(), s.B.Bytes()).BigInt()

	tmp := big.NewInt(0).Mul(A, big.NewInt(0).Exp(s.v, u, s.N))
	secret := big.NewInt(0).Exp(tmp, s.b, s.N)

	// fmt.Println("A*(v^u mod N) length ", len(tmp.Bytes()) * 8)
	// fmt.Println("A*(v^u mod N) = ", codec.ToByteArray(tmp.Bytes()).Base64())
	// fmt.Println("(A*(v^u mod N)) ^ b mod N = ", len(secret.Bytes()) * 8)

	s.K = codec.GetHash(codec.SHA256, secret.Bytes())
	s.M = cryptowrapper.HMACSHA256(s.K, A.Bytes(), s.B.Bytes(), s.e, s.S, s.N.Bytes(), s.g.Bytes())

	return
}

// CheckClient authenticates the client
func (s *SRPServer) CheckClient(auth []byte) (proof []byte, err error) {
	if bytes.Equal(s.M, auth) {
		proof = cryptowrapper.HMACSHA256(s.K, auth)
		return proof, err
	}
	return []byte(""), errors.New("client authentication failed")
}

package srp

import (
	"errors"
	"math/big"
	"testing"

	"github.com/TectusDreamlab/go-common-utils/cryptowrapper"
	. "github.com/smartystreets/goconvey/convey"
)

func TestSRP(t *testing.T) {
	Convey("Test Secure Remote Passwrod Implementation", t, func() {
		Convey("Create Client With Wrong Bit Length Should Fail", func() {
			_, err := NewSRPClient("test@gmail.com", "testpassword", 1234)
			So(err, ShouldBeError, errors.New("invalid bits"))
		})

		client, err := NewSRPClient("test@gmail.com", "testpassword", 4096)
		client.GetX()
		verifier := client.GetVerifier()
		Convey("Create Client With Length 4096 Should Be Successful", func() {
			So(err, ShouldBeNil)
			So(verifier, ShouldNotBeNil)
		})

		Convey("Create Server With Wrong Bit Length Should Fail", func() {
			_, err := NewSRPServer(client.E, client.S, client.GetVerifier(), client.A, 1234)
			So(err, ShouldBeError, errors.New("invalid bits"))
		})

		server, err := NewSRPServer(client.E, client.S, verifier, client.A, 4096)
		Convey("Create Server With Length 4096 And Corresponding Client Verifier, Salt And Public A Should Be Successful", func() {
			So(err, ShouldBeNil)
		})

		Convey("Server And Client Should Calculate The Same K And Be Able To Check Each Other", func() {
			err := client.Calculate(server.B, server.S)
			So(err, ShouldBeNil)
			So(server.K, ShouldResemble, client.K)
			proof, err := server.CheckClient(client.M)
			So(err, ShouldBeNil)
			err = client.CheckServer(proof)
			So(err, ShouldBeNil)
		})

		Convey("Test Error Cases", func() {
			Convey("Client Gets A Wrong Public B From Server Should Cause An Error", func() {
				err = client.Calculate(big.NewInt(0), server.S)
				So(err, ShouldBeError, errors.New("invalid server public key B=0"))
			})

			Convey("Server Did Not Receive The Corrent Verifier Should Cause Mutual Check Failed", func() {
				server, err := NewSRPServer(client.E, client.S, cryptowrapper.RandBigInt(128), client.A, 4096)
				So(err, ShouldBeNil)
				err = client.Calculate(server.B, server.S)
				So(err, ShouldBeNil)
				proof, err := server.CheckClient(client.M)
				So(err, ShouldBeError, errors.New("client authentication failed"))
				err = client.CheckServer(proof)
				So(err, ShouldBeError, errors.New("server authentication failed"))
			})

			Convey("Server Did Not Receive The Corrent Public A Should Cause Mutual Check Failed", func() {
				server, err := NewSRPServer(client.E, client.S, verifier, cryptowrapper.RandBigInt(128), 4096)
				So(err, ShouldBeNil)
				err = client.Calculate(server.B, server.S)
				So(err, ShouldBeNil)
				proof, err := server.CheckClient(client.M)
				So(err, ShouldBeError, errors.New("client authentication failed"))
				err = client.CheckServer(proof)
				So(err, ShouldBeError, errors.New("server authentication failed"))
			})

			Convey("Client Did Not Receive The Corrent Public B Should Cause Mutual Check Failed", func() {
				err = client.Calculate(cryptowrapper.RandBigInt(128), server.S)
				So(err, ShouldBeNil)
				proof, err := server.CheckClient(client.M)
				So(err, ShouldBeError, errors.New("client authentication failed"))
				err = client.CheckServer(proof)
				So(err, ShouldBeError, errors.New("server authentication failed"))
			})
		})
	})
}

package crypto

/*
import (
	"io"
	"log"

	tpm2tools "github.com/google/go-tpm-tools/client"
	"github.com/google/go-tpm/tpm2"
)

func openTPM() (rwc io.ReadWriteCloser, err error) {
	rwc, err = tpm2.OpenTPM(tpmPath)
	if err != nil {
		log.Fatalf("can't open TPM %q: %v", tpmPath, err)
	}
	return
}

func getTPMPublicKey(rwc io.ReadWriteCloser) (srk, err error) {
	srk, err = tpm2tools.StorageRootKeyRSA(rwc)
	if err != nil {
		log.Fatalf("can't create srk from template: %v", err)
	}
	return
}

func Seal(secret []byte, srk io.ReadWriteCloser) (sealed []byte, err error) {
	rwc, err := openTPM()
	defer func() {
		if err := rwc.Close(); err != nil {
			log.Fatalf("%v\ncan't close TPM %q: %v", retErr, tpmPath, err)
		}
	}()
	if err != nil {
		//
		return
	}

	srk, err := getTPMPublicKey(rwc)
	defer srk.Close()
	if err != nil {
		//
		return
	}

	//sel := tpm2.PCRSelection{Hash: tpm2.AlgSHA256, PCRs: []int{pcr}}
	//sOpt := tpm2tools.SealCurrent{PCRSelection: sel}

	sealed, err := srk.Seal(secret, nil) // sOPT
	if err != nil {
		log.Fatalf("failed to seal: %v", err)
	}
	return sealed, err
}

func Unseal(sealedSecret *tpmpb.SealedBytes) (unsealedSecret []byte, err error) {
	rwc, err := openTPM()
	defer func() {
		if err := rwc.Close(); err != nil {
			log.Fatalf("%v\ncan't close TPM %q: %v", retErr, tpmPath, err)
		}
	}()
	if err != nil {
		//
		return
	}

	srk, err := getTPMPublicKey(rwc)
	defer srk.Close()
	if err != nil {
		//
		return
	}

	unsealedSecret, err := srk.Unseal(sealedSecret, nil) // cOPT to verify state of TPM
	if err != nil {
		log.Fatalf("failed to seal: %v", err)
	}
	return sealed, err
}

*/

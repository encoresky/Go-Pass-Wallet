package internal

import (
	"net/http"
	"time"

	"github.com/alvinbaena/passkit"
	"golang.org/x/exp/slog"
)

func GeneratePass(w http.ResponseWriter, r *http.Request) {

	newEvent := passkit.NewEventTicket()
	newEvent.AddHeaderField(passkit.Field{
		Key:           "header_key",
		Label:         "HEADER",
		Value:         "header value",
		TextAlignment: passkit.TextAlignmentRight,
	})

	newEvent.AddSecondaryFields(passkit.Field{
		Key:           "secondary_key",
		Label:         "SECONDARY",
		Value:         "secondary value",
		TextAlignment: passkit.TextAlignmentLeft,
	})

	newEvent.AddAuxiliaryFields(passkit.Field{
		Key:           "auxiliary_key1",
		Label:         "AUXIILIARY1",
		Value:         "auxiliary_value",
		TextAlignment: passkit.TextAlignmentLeft,
	})

	newEvent.AddAuxiliaryFields(passkit.Field{
		Key:           "auxiliary_key2",
		Label:         "AUXIILIARY2",
		Value:         "auxiliary_value",
		TextAlignment: passkit.TextAlignmentLeft,
	})

	newEvent.AddAuxiliaryFields(passkit.Field{
		Key:           "auxiliary_key3",
		Label:         "AUXIILIARY3",
		Value:         "auxiliary_value",
		TextAlignment: passkit.TextAlignmentLeft,
	})

	pkPass := passkit.Pass{
		WebServiceURL:       "http:\\localhost:3001", // replace with your application url
		FormatVersion:       1.0,
		TeamIdentifier:      "7BB29FUWZL", // replace by team id
		PassTypeIdentifier:  "XYZ-123",
		AuthenticationToken: "PASS-XYZ-ABCalkaaksjlasasdfjaskjfalsdfjaksdjflaskdjflasdkfjlasdkjflasdkfjasdkjflasdkfj",
		OrganizationName:    "orgnization name", // replace by your orgnization name
		SerialNumber:        "1234567890",
		Description:         "description", // write your organization description
		LabelColor:          "#FFFFFF",
		ForegroundColor:     "#FFEDD7",
		BackgroundColor:     "#006F7D",
		EventTicket:         newEvent,
		Barcodes: []passkit.Barcode{
			{
				Format:          passkit.BarcodeFormatQR,
				Message:         "http:\\localhost:3001",
				MessageEncoding: "iso-8859-1",
			},
		},
	}

	template := passkit.NewInMemoryPassTemplate()

	template.AddFileBytes(passkit.BundleLogo, []byte(`replace this with image bytes`))

	template.AddFileBytes(passkit.BundleIcon, []byte(`replace this with image bytes`))

	signInfo, err := passkit.LoadSigningInformationFromFiles("./internal/certificate/Certificates.p12", "pass@2022", "./internal/certificate/AppleWWDRCAG4.cer")
	if err != nil {
		slog.Error("cannot load singing information from files", "error", err)
		http.Error(w, "cannot load singing information from files", http.StatusInternalServerError)
		return
	}

	signer := passkit.NewMemoryBasedSigner()
	pkpassBytes, err := signer.CreateSignedAndZippedPassArchive(&pkPass, template, signInfo)
	if err != nil {
		slog.Error("cannot signed and zipped pass", "error", err)
		http.Error(w, "cannot signed and zipped pass", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename=demo.pkpass")
	w.Header().Set("Content-Type", "application/vnd.apple.pkpass")
	w.Header().Set("Last-Modified", time.Now().GoString())
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(pkpassBytes)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}

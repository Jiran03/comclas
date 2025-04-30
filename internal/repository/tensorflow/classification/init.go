package classificationrepository

import (
	"fmt"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

func Classification() {
	model, err := tf.LoadSavedModel("comment_classification_model", []string{"serve"}, nil)
	if err != nil {
		panic(fmt.Sprintf("Gagal memuat model: %v", err))
	}

	defer model.Session.Close()
}

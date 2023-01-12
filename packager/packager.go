package packager

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/dchest/uniuri"
	"io"
	"os"
)

func Packager(fileInput string) {
	// Main File Info
	filemeta := FileHeader{
		Magic:   "NTXPCK",
		Version: 2,
	}
	//Open Input File
	inFile, err := os.Open(fileInput)
	if err != nil {
		panic(err)
	}
	//Create Random Key
	key := uniuri.NewLen(16)
	filemeta.Key = []byte(key)

	//Create Encryption Block
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	iv := make([]byte, block.BlockSize())
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	filemeta.iv = iv
	outFile, err := os.Create(inFile.Name() + ".pck")
	if err != nil {
		panic(err)
	}
	// ALWAYS TO CLOSE IN THE END OF THE PROCESS
	defer outFile.Close()
	defer inFile.Close()

	//Make Chunk for 1024 Bytes
	buffer := make([]byte, 1024)
	//CTR Stream
	stream := cipher.NewCTR(block, iv)

	//Write Metadata Header
	outFile.Write([]byte(filemeta.Magic))   //MAGIC CODE
	outFile.Write([]byte{filemeta.Version}) //Version
	outFile.Write(filemeta.Key)             //Key
	outFile.Write(filemeta.iv)              //IV

	//Read All Byte until the end of the File
	for {
		numberOfBytes, err := inFile.Read(buffer)
		if numberOfBytes > 0 {
			stream.XORKeyStream(buffer, buffer[:numberOfBytes])
			//Write into File
			outFile.Write(buffer[:numberOfBytes])
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Failed to Read Bytes: %v", err)
			panic(err)
		}
	}
	//Done
	fmt.Printf("File is written to %s.pck", inFile.Name())
}

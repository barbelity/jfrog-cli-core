package commands

import (
    "github.com/jfrogdev/jfrog-cli-go/utils/cliutils"
    "github.com/jfrogdev/jfrog-cli-go/bintray/utils"
    "github.com/jfrogdev/jfrog-cli-go/bintray/tests"
    "strconv"
	"testing"
)

func TestSingleFileUpload(t *testing.T) {
    versionDetails := utils.CreateVersionDetails("test-subject/test-repo/test-package/ver-1.2")
    uploadPath := "/a/b/"
    flags := createUploadFlags()

	uploaded1, _ := Upload(versionDetails, "testdata/a.txt", uploadPath, flags)
	uploaded2, _ := Upload(versionDetails, "testdata/aa.txt", uploadPath, flags)
	uploaded3, _ := Upload(versionDetails, "testdata/aa1*.txt", uploadPath, flags)

	if uploaded1 != 1 {
		t.Error("Expected 1 file to be uploaded. Got " + strconv.Itoa(uploaded1) + ".")
	}
	if uploaded2 != 1 {
		t.Error("Expected 1 file to be uploaded. Got " + strconv.Itoa(uploaded2) + ".")
	}
	if uploaded3 != 0 {
		t.Error("Expected 0 file to be uploaded. Got " + strconv.Itoa(uploaded3) + ".")
	}
}

func TestPatternRecursiveUpload(t *testing.T) {
	flags := createUploadFlags()
	flags.Recursive = true
	testPatternUpload(t, flags)
}

func TestPatternNonRecursiveUpload(t *testing.T) {
	flags := createUploadFlags()
	flags.Recursive = false
	testPatternUpload(t, flags)
}

func testPatternUpload(t *testing.T, flags *UploadFlags) {
    versionDetails := utils.CreateVersionDetails("test-subject/test-repo/test-package/ver-1.2")
    uploadPath := "/a/b/"

	sep := cliutils.GetTestsFileSeperator()
	uploaded1, _ := Upload(versionDetails, "testdata"+sep+"*", uploadPath, flags)
	uploaded2, _ := Upload(versionDetails, "testdata"+sep+"a*", uploadPath, flags)
	uploaded3, _ := Upload(versionDetails, "testdata"+sep+"b*", uploadPath, flags)

	if uploaded1 != 3 {
		t.Error("Expected 3 file to be uploaded. Got " + strconv.Itoa(uploaded1) + ".")
	}
	if uploaded2 != 2 {
		t.Error("Expected 2 file to be uploaded. Got " + strconv.Itoa(uploaded2) + ".")
	}
	if uploaded3 != 1 {
		t.Error("Expected 1 file to be uploaded. Got " + strconv.Itoa(uploaded3) + ".")
	}
}

func createUploadFlags() *UploadFlags {
	return &UploadFlags{
		BintrayDetails: tests.CreateBintrayDetails(),
		Recursive:      true,
		Flat:           true,
		Publish:        false,
		Override:       false,
		Explode:        false,
		UseRegExp:      false,
		Threads:        3,
		Deb:            "",
		DryRun:         true}
}
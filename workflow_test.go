package app

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/testsuite"
)

func Test_Workflow(t *testing.T) {
	// Set up the test suite and testing execution environment
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	// Mock activity implementation
	env.OnActivity(TryDownload, mock.Anything, "https://download.library.lol/main/4154000/387eb49ed942fda6a0ef896a98ecceae/%28Literature%20Now%29%20Dan%20Sinykin%20-%20Big%20Fiction_%20How%20Conglomeration%20Changed%20the%20Publishing%20Industry%20and%20American%20Literature-Columbia%20University%20Press%20%282023%29.pdf", "/tmp").Return("file is 3225167 bytes\n", nil)

	env.ExecuteWorkflow(DownloadWorkflow, "https://download.library.lol/main/4154000/387eb49ed942fda6a0ef896a98ecceae/%28Literature%20Now%29%20Dan%20Sinykin%20-%20Big%20Fiction_%20How%20Conglomeration%20Changed%20the%20Publishing%20Industry%20and%20American%20Literature-Columbia%20University%20Press%20%282023%29.pdf", "/tmp")
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var size string
	require.NoError(t, env.GetWorkflowResult(&size))
	require.Equal(t, "file is 3225167 bytes\n", size)
}

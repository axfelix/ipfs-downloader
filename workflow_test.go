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
	env.OnActivity(TryDownload, mock.Anything, "https://cloudflare-ipfs.com/ipfs/bafykbzacecdmotlibikawsyygjqrlzz6twacvcsndhhpps52t3hzhoqoaqpzy?filename=Big%20Fiction%3A%20How%20Conglomeration%20Changed%20the%20Publishing%20--%20Dan%20Sinykin%20--%20Literature%20Now%2C%202023%20--%20Columbia%20University%20Press%20--%209780231192941%20--%20387eb49ed942fda6a0ef896a98ecceae%20--%20Anna%E2%80%99s%20Archive.pdf", "~/Desktop").Return("", nil)

	env.ExecuteWorkflow(GreetingWorkflow, "https://cloudflare-ipfs.com/ipfs/bafykbzacecdmotlibikawsyygjqrlzz6twacvcsndhhpps52t3hzhoqoaqpzy?filename=Big%20Fiction%3A%20How%20Conglomeration%20Changed%20the%20Publishing%20--%20Dan%20Sinykin%20--%20Literature%20Now%2C%202023%20--%20Columbia%20University%20Press%20--%209780231192941%20--%20387eb49ed942fda6a0ef896a98ecceae%20--%20Anna%E2%80%99s%20Archive.pdf", "~/Desktop")
	require.True(t, env.IsWorkflowCompleted())
	require.NoError(t, env.GetWorkflowError())

	var size string
	require.NoError(t, env.GetWorkflowResult(&size))
	require.Equal(t, "Hello World!", size)
}

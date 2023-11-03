package v1alpha1

type Operation struct {
	// Assert represents the assertions to be made for this test step. It checks whether the conditions
	// specified in each assertion hold true.
	// +optional
	Assert []Assert `json:"assert,omitempty"`

	// Apply lists the resources that should be applied for this test step. This can include things
	// like configuration settings or any other resources that need to be available during the test.
	// +optional
	Apply []Apply `json:"apply,omitempty"`

	// Create represents a creation operation.
	// +optional
	Create *Create `json:"create,omitempty"`

	// Error lists the expected errors for this test step. If any of these errors occur, the test
	// will consider them as expected; otherwise, they will be treated as test failures.
	// +optional
	Error []Error `json:"error,omitempty"`

	// Delete provides a list of objects that should be deleted before this test step is executed.
	// This helps in ensuring that the environment is set up correctly before the test step runs.
	// +optional
	Delete []Delete `json:"delete,omitempty"`

	// Exec provides a list of commands and/or scripts that should be executed as a part of this test step.
	// +optional
	Exec []ExecOperation `json:"exec,omitempty"`
}
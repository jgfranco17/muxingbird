package errorx

// ShellExitCode represents a CLI process exit code.
type ShellExitCode int

const (
	// ExitSuccess indicates successful execution.
	ExitSuccess ShellExitCode = 0

	// ExitGenericError is a non-specific error.
	ExitGenericError ShellExitCode = 1

	// ExitInvalidArgs is returned when CLI args are invalid.
	ExitInvalidArgs ShellExitCode = 2

	// ExitFileNotFound indicates a missing input file.
	ExitFileNotFound ShellExitCode = 3

	// ExitConfigError represents a config parsing error.
	ExitConfigError ShellExitCode = 4

	// ExitServerError represents internal server launch/exec errors.
	ExitServerError ShellExitCode = 5

	// ExitPanic indicates an unhandled panic or fatal error.
	ExitPanic ShellExitCode = 100
)

package database

// DBError is used when unspecific database operations fail / rows.Scan fails
type DBError struct {
    Message string
    InnerErr error
}

func (err DBError) Error() string {
    if err.InnerErr != nil {
        return err.Message + ": " + err.InnerErr.Error()
    }
    return err.Message
}

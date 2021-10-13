package database

// DBError is used when unspecific database operations fail / rows.Scan fails
type InternalServerError struct {
    Message string
    InnerErr error
}

type BadRequest struct {
    Message string
    InnerErr error
}

func (err InternalServerError) Error() string {
    if err.InnerErr != nil {
        return err.Message + ": " + err.InnerErr.Error()
    }
    return err.Message
}

func (err BadRequest) Error() string {
    if err.InnerErr != nil {
        return err.Message + ": " + err.InnerErr.Error()
    }
    return err.Message
}

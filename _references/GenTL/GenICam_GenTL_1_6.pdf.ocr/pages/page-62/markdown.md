|  ![img-91.jpeg](img-91.jpeg) CAM |   | ![img-92.jpeg](img-92.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|  GC_ERR_ACCESS_DENIED | -1005 | Requested operation is not allowed; e.g., a remote device is opened by another client.  |
|  GC_ERR_INVALID_HANDLE | -1006 | Given handle does not support the operation; e.g., function call on wrong handle or NULL pointer.  |
|  GC_ERR_INVALID_ID | -1007 | ID could not be connected to a resource; e.g., a device with the given ID is currently not available.  |
|  GC_ERR_NO_DATA | -1008 | The function has no data to work on or the data does not provide reliable information corresponding with the request.  |
|  GC_ERR_INVALID_PARAMETER | -1009 | One of the parameter given was not valid or out of range.  |
|  GC_ERR_IO | -1010 | Communication error has occurred; e.g., a read or write operation to a remote device failed.  |
|  GC_ERR_TIMEOUT | -1011 | An operation's timeout time expired before it could be completed.  |
|  GC_ERR_ABORT | -1012 | An operation has been aborted before it could be completed. For example a wait operation through EventGetData has been terminated via a call to EventKill.  |
|  GC_ERR_INVALID_BUFFER | -1013 | The GenTL Consumer has not announced enough buffers to start the acquisition in the currently active acquisition mode.  |
|  GC_ERR_NOT_AVAILABLE | -1014 | Resource or information is not available at a given time in a current state.  |
|  GC_ERR_INVALID_ADDRESS | -1015 | A given address is out of range or invalid for internal reasons.  |
|  GC_ERR_BUFFER_TOO_SMALL | -1016 | A provided buffer is too small to receive the expected amount of data. This may affect acquisition buffers in the Data Stream module if the buffers are smaller than the expected payload size but also buffers passed to any other function of the GenTL Producer interface to retrieve information or IDs.  |
|  GC_ERR_INVALID_INDEX | -1017 | A provided index referencing a  |
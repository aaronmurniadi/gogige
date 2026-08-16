|  Category | FileAccessControl  |
| --- | --- |
|  Level | Recommended  |
|  Interface | Integer  |
|  Access | Read/Write  |
|  Unit | B  |
|  Visibility | Guru  |
|  Values | ≥ 0  |

Controls the Length of the mapping between the device file storage and the FileAccessBuffer.

The FileAccessLength defines the number of bytes to transfer to or from the FileAccessBuffer (See Figure 18-2). This feature is available only when FileOperationSelector is set to Read or Write.

### 18.9 FileOperationStatus

|  Name | FileOperationStatus[FileSelector][FileOperationSelector]  |
| --- | --- |
|  Category | FileAccessControl  |
|  Level | Recommended  |
|  Interface | Enumeration  |
|  Access | Read  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Success Failure ...  |

Represents the file operation execution status.

Upon execution of a successful file operation, it must return Success. In case of complete or partial failure of the operation, other return values can be defined to indicate the nature of the error that happened. If only one fail status is defined, it should be defined as Failure.

Possible values are:

- Success: File Operation was successful.
- Failure: File Operation failed.

### 18.10 FileOperationResult

|  Name | FileOperationResult[FileSelector][FileOperationSelector]  |
| --- | --- |
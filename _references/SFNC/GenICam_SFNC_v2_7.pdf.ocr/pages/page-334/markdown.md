### 18.6 FileAccessBuffer

|  Name | FileAccessBuffer  |
| --- | --- |
|  Category | FileAccessControl  |
|  Level | Recommended  |
|  Interface | IRegister  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Device-specific  |

Defines the intermediate access buffer that allows the exchange of data between the device file storage and the application.

This register mapped FileAccessBuffer must be written with the target data before executing a Write operation. For Read Operation, the data can be read from the FileAccessBuffer after the Read operation has been executed. The effective data transfer is done upon FileOperationExecute execution (See Figure 18-2: Layout of File Access Buffer.).

### 18.7 FileAccessOffset

|  Name | FileAccessOffset[FileSelector][FileOperationSelector]  |
| --- | --- |
|  Category | FileAccessControl  |
|  Level | Recommended  |
|  Interface | IInteger  |
|  Access | Read/(Write)  |
|  Unit | B  |
|  Visibility | Guru  |
|  Values | ≥ 0  |

Controls the Offset of the mapping between the device file storage and the FileAccessBuffer.

The FileAccessOffset defines the offset in bytes of the FileAccessBuffer relative to the beginning of the selected File (See Figure 18-2). This feature is available only when FileOperationSelector is set to Read or Write.

### 18.8 FileAccessLength

|  Name | FileAccessLength[FileSelector][FileOperationSelector]  |
| --- | --- |
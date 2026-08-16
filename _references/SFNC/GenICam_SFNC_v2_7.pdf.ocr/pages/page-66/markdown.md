![img-8.jpeg](img-8.jpeg)

Table 2-16: File Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  FileAccessControl | R | ICategory | R | - | G | Category that contains the File Access control features.  |
|  FileSelector | R | IEnumeration | R/(W) | - | G | Selects the target file in the device.  |
|  FileOperationSelector[FileSelector] | R | IEnumeration | R/W | - | G | Selects the target operation for the selected file in the device.  |
|  FileOperationExecute[FileSelector][FileOperationSelector] | R | ICommand | (R)/W | - | G | Executes the operation selected by FileOperationSelector on the selected file.  |
|  FileOpenMode[FileSelector] | R | IEnumeration | R/(W) | - | G | Selects the access mode in which a file is opened in the device.  |
|  FileAccessBuffer | R | IRegister | R/(W) | - | G | Defines the intermediate access buffer that allows the exchange of data between the device file storage and the application.  |
|  FileAccessOffset[FileSelector][FileOperationSelector] | R | IInteger | R/(W) | B | G | Controls the Offset of the mapping between the device file storage and the FileAccessBuffer.  |
|  FileAccessLength[FileSelector][FileOperationSelector] | R | IInteger | R/W | B | G | Controls the Length of the mapping between the device file storage and the FileAccessBuffer.  |
|  FileOperationStatus[FileSelector][FileOperationSelector] | R | IEnumeration | R | - | G | Represents the file operation execution status.  |
|  FileOperationResult[FileSelector][FileOperationSelector] | R | IInteger | R | - | G | Represents the file operation result.  |
|  FileSize[FileSelector] | R | IInteger | R | B | G | Represents the size of the selected file in bytes.  |

## 2.17 Source Control

Contains the features related to the control of the multiple Source devices (See the Source Control chapter for details).

Table 2-17: Source Control Summary

|  Name | Level | Interface | Access | Unit | Visibility | Description  |
| --- | --- | --- | --- | --- | --- | --- |
|  SourceControl | O | ICategory | R | - | B | Category that contains the source control features.  |
|  SourceCount | O | IInteger | R/(W) | - | B | Controls or returns the number of sources supported by the device.  |
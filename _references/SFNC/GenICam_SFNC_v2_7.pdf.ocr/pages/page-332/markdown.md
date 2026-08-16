- **UserSetDefault**: The default user set of the device.
- **UserSet1**: The first user set of the device.
- **UserSet2**: The second user set of the device.
- **UserSet3**: The third user set of the device.
- ...
- **LUTLuminance**: The Luminance LUT of the camera.
- **LUTRed**: The Red LUT of the camera.
- **LUTGreen**: The Green LUT of the camera.
- **LUTBlue**: The Blue LUT of the camera.
- ...

On top of the previous standard values, a device might also provide device-specific values.

### 18.3 FileOperationSelector

|  Name | FileOperationSelector[FileSelector]  |
| --- | --- |
|  Category | FileAccessControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Open Close Read Write Delete  |

Selects the target operation for the selected file in the device. This Operation is executed when the **FileOperationExecute** feature is called.

Possible values are:

- **Open**: Opens the file selected by **FileSelector** in the device. The access mode in which the file is opened is selected by **FileOpenMode**.
- **Close**: Closes the file selected by **FileSelector** in the device.
- **Read**: Reads **FileAccessLength** bytes from the device storage at the file relative offset **FileAccessOffset** into **FileAccessBuffer**.
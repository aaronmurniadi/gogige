|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- Write: Writes FileAccessLength bytes taken from the FileAccessBuffer into the device storage at the file relative offset FileAccessOffset.
- Delete: Deletes the file selected by FileSelector in the device. Note that deleting a device file should not remove the associated FileSelector entry to allow future operation on this file.

### 18.4 FileOperationExecute

|  Name | FileOperationExecute[FileSelector][FileOperationSelector]  |
| --- | --- |
|  Category | FileAccessControl  |
|  Level | Recommended  |
|  Interface | ICommand  |
|  Access | (Read)/Write  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | -  |

Executes the operation selected by FileOperationSelector on the selected file.

### 18.5 FileOpenMode

|  Name | FileOpenMode[FileSelector]  |
| --- | --- |
|  Category | FileAccessControl  |
|  Level | Recommended  |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Guru  |
|  Values | Read Write ReadWrite  |

Selects the access mode in which a file is opened in the device.

Possible values are:

- Read: This mode selects read-only open mode.
- Write: This mode selects write-only open mode.
- ReadWrite: This mode selects read and write open mode.
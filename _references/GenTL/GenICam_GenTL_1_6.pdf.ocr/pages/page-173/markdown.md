|  ![img-256.jpeg](img-256.jpeg) CAM |   | ![img-257.jpeg](img-257.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

|  Enumerator | Value | Description  |
| --- | --- | --- |
|   |  | EVENT_DATA_IDData type: INT32 (GC_ERROR)EVENT_DATA_VALUEData type: STRING (Description)  |
|  EVENT_NEW_BUFFER | 1 | Notification on newly filled buffers. Values that can be retrieved are:EVENT_DATA_IDData type: PTR (Buffer handle)EVENT_DATA_VALUEData type: PTR (Private pointer)  |
|  EVENT_FEATURE_INVALIDATE | 2 | Notification if a feature was changed by the GenTL Producer driver and thus needs to be invalidated in the GenICam GenApi instance using the module. Values that can be retrieved are:EVENT_DATA_IDData type: STRING (Feature name)  |
|  EVENT_FEATURE_CHANGE | 3 | Notification if the GenTL Producer driver wants to manually set a feature in the GenICam GenApi instance using the module. Values that can be retrieved are:EVENT_DATA_IDData type: STRING (Feature name)EVENT_DATA_VALUEData type: STRING (Feature value)  |
|  EVENT_REMOTE_DEVICE | 4 | Notification if the GenTL Producer wants to inform the GenICam GenApi instance of the remote device that a GenApi compatible event was fired. This Event is to be registered on a Local Device module.Values that can be retrieved are:EVENT_DATA_IDString representation of the EventID number in hexadecimal numbers with even number of digits and without the leading ‘0x’.Data type: STRING (Event ID)EVENT_DATA_VALUECorresponds to the data addressable through the remote device's nodemap event port, beginning of the buffer corresponding to address 0.Data type: BUFFER (optional data)  |
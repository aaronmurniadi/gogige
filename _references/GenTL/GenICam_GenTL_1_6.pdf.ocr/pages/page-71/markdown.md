|  ![img-99.jpeg](img-99.jpeg) CAM |   | ![img-100.jpeg](img-100.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.2.4 TLGetInterfaceInfo

GC_ERROR TLGetInterfaceInfo ( TL_HANDLE hSystem,
    const char * sIfaceID,
    INTERFACE_INFO_CMD iInfoCmd,
    INFO_DATATYPE * piType,
    void * pBuffer,
    size_t * piSize )

Inquire information about an interface on the given System module hSystem as defined in INTERFACE_INFO_CMD without opening the interface. The reported information should be in sync to information retrieved through the IFGetInfo function.

Parameters

|  [in] | hSystem | System module to work on.  |
| --- | --- | --- |
|  [in] | sIfaceID | Unique ID of the interface to inquire information from. Like with theTLOpenInterfacefunction it is also possible to feed an alternative ID as long as the GenTL Producer knows how to interpret it.  |
|  [in] | iInfoCmd | Information to be retrieved as defined inINTERFACE_INFO_CMD.  |
|  [out] | piType | Data type of thepBuffercontent as defined in theINTERFACE_INFO_CMDandINFO_DATATYPE.  |
|  [in,out] | pBuffer | Pointer to a user allocated buffer to receive the requested information. If this parameter isNULL,piSizewill contain the minimal size ofpBufferin bytes. If thepiTypeis a string the size includes the terminating 0.  |
|  [in,out] | piSize | pBufferequal NULL:out: minimal size ofpBufferin bytes to hold all informationpBufferunequal NULL:in: size of the providedpBufferin bytesout: number of bytes filled by the function  |

Returns

|  GC_ERR_SUCCESS | Operation was successful; no error occurred.  |
| --- | --- |
|  GC_ERR_NOT_INITIALIZED | No preceding call to GCInitLib.  |
|  GC_ERR_INVALID_HANDLE | The handle hSystem is invalid (NULL) or does not reference an open System module retrieved through a call to TLOpen.  |
|  GC_ERR_NOT_IMPLEMENTED | Specified iInfoCmd is not implemented.  |
|  GC_ERR_INVALID_PARAMETER | Parameters piSize and/or piType are invalid pointers (NULL or ~0x0).  |
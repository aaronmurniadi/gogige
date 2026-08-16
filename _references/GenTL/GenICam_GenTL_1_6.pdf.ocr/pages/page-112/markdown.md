|  ![img-152.jpeg](img-152.jpeg) CAM |   | ![img-153.jpeg](img-153.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

##### 6.3.5.22 DSGetBufferSegmentInfo

|  GC_ERROR | DSGetBufferSegmentInfo ( DS_HANDLE hDataStream, BUFFER_HANDLE hBuffer, uint32_t iSegmentIndex, SEGMENT_INFO_CMD iInfoCmd, INFO_DATATYPE * piType, void * pBuffer, size_t * piSize )  |
| --- | --- |

Inquires information about individual segments of a buffer as defined in SEGMENT_INFO_CMD.

If hBuffer references a composite buffer (announced using DSAnnounceCompositeBuffer), the function reports information about “real” segments, as announced in DSAnnounceCompositeBuffer.

If hBuffer references a contiguous buffer (announced using DSAllocAndAnnounceBuffer or DSAnnounceBuffer), the function reports information about virtual segments created within the buffer by the acquisition engine when it was last time filled. When flows were used for the acquisition, the virtual segments correspond to the flow structure  \( (5.7.2) \) , otherwise the buffer is assumed to consist of a single segment.

##### Parameters

|  [in] | hDataStream | Data Stream module to work on.  |
| --- | --- | --- |
|  [in] | hBuffer | Buffer handle to retrieve information about.  |
|  [in] | iSegmentIndex | Zero based index of the buffer segment to query.  |
|  [in] | iInfoCmd | Information to be retrieved as defined in SEGMENT_INFO_CMD.  |
|  [out] | piType | Data type of the pBuffer content as defined in the SEGMENT_INFO_CMD and INFO_DATATYPE.  |
|  [in,out] | pBuffer | Pointer to a user allocated buffer to receive the requested information. If this parameter is NULL, piSize will contain the minimal size of pBuffer in bytes. If the piType is a string the size includes the terminating 0.  |
|  [in,out] | piSize | pBuffer equal NULL: out: minimal size of pBuffer in bytes to hold all information. pBuffer unequal NULL: in: size of the provided pBuffer in bytes. out: number of bytes filled by the function.  |

##### Returns

GC_ERR_SUCCESS: Operation was successful; no error occurred.

GC_ERR_NOT_INITIALIZED: No preceding call to GCInitLib.
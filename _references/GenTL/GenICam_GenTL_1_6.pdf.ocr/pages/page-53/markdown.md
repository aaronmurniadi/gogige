|  ![img-72.jpeg](img-72.jpeg) CAM |   | ![img-73.jpeg](img-73.jpeg) emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

BUFFER_INFO_WIDTH, BUFFER_INFO_HEIGHT, etc.). Even if some of the properties, such as AOI size, are the same for all parts in a buffer, it should be reported and queried per-part via DSGetBufferPartInfo, using the BUFFER_PART_INFO_CMDs.

On the other hand buffer properties which are describing the global buffer and are not defined as part-specific have to be queried using DSGetBufferInfo (e.g., BUFFER_INFO_TIMESTAMP, BUFFER_INFO_NEW_DATA, BUFFER_INFO_DELIVERED_CHUNKPAYLOADSIZE, etc). It is listed with the BUFFER_INFO_XXX constants which constant is overwritten by part specific information or if it describes the whole buffer.

Similar to any other basic payload type, it is possible to attach chunk data to the multi-part payload. The principles of chunk data handling remain the same as with other basic payload types. The chunk data is therefore common to all parts in the buffer. There is only one chunk section within a multi-part buffer.

The GenTL specification does not define strict rules for relationships between the individual data types within a buffer. Some typical use cases are discussed in the text below.

#### 5.6.2 Planar Pixel Formats

A multi-part buffer can be used to reliably transfer and describe data using a planar pixel format such as the color data in separate R-G-B planes. In this case each part carries a single color plane. Typically all the parts share the same dimensions and differ only in the data format. For example, in this case the used data type would be (depending on the actual data) PART_DATATYPE_2D_PLANE_TRIPLANAR.

With multi-part buffer all the planes that belong together are well and unambiguously described.

#### 5.6.3 Multiple AOI's

There are devices which support multiple areas of interest (AOI's) to be captured within the sensor. The data from these multiple AOI's can also be effectively transferred using the multi-part payload. In this case the data format of all the parts is typically the same, the parts differ only in the AOI parameters.

It is up to the GenTL Consumer if it will prefer to treat the individual AOI's as independent images or if it will reconstruct a single image from the AOI's. The main advantage is that all the AOI's belonging to the same exposure are transferred together.

#### 5.6.4 Pixel Confidence Data

A buffer part carrying the data type PART_DATATYPE CONFIDENCE_MAP is used to identify levels of validity of the pixel values carried in other part(s). Each value in the confidence map specifies the confidence level of pixels at the same position (row/column) in other data part(s).

In the simplest case of a 1-bit confidence data the confidence map simply marks corresponding pixels valid or invalid. Higher bit depth integer data types in the validity map allow to specify the level of confidence ranging from 0 to the maximal value of given integer
|  GEN<i>CAM |   | ![img-25.jpeg](img-25.jpeg) emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

|  Level | Optional  |
| --- | --- |
|  Interface | IEnumeration  |
|  Access | Read/(Write)  |
|  Unit | -  |
|  Visibility | Beginner  |
|  Values | Region0 (if 0 based)Region1Region2...Scan3dExtraction0 (if 0 based)Scan3dExtraction1Scan3dExtraction2...All  |

Selects the Region of interest to control. The RegionSelector feature allows devices that are able to extract multiple regions out of an image, to configure the features of those individual regions independently.

RegionX are generally used to configure Sensor's regions extraction characteristics. Regions can overlap and can also be used to provide the data of the same region in various PixelFormat or data compression type.

For example to provide the RGB Uncompressed Intensity and the YUV JPEG compressed Intensity of a sensor region:

RegionSelector = Region0 // Select Region 0
RegionMode[Region0] = On // Instance 0 (default ON)
ComponentSelector = Intensity // Select Intensity
ComponentEnable[Region0][Intensity] = true // Enable instance 0 streaming (default).
PixelFormat[Region0][Intensity] = RGB8 // RGB 24 bit per pixel.
ImageCompressionMode[Region0][Intensity] = Off // Instance 0 Uncompressed.

RegionSelector = Region1 // Select Region 1
RegionMode[Region1] = On // Instance 1
ComponentSelector = Intensity // Select Intensity
Component [Intensity] = true // Enable instance 1 streaming.
PixelFormat[Region1][Intensity] = YUV422_8 // YUV 16 bit per pixel.
ImageCompressionMode[Region1][Intensity] = JPEG // Instance 1 Compressed

Other Processing module output Regions can be used to configure the size and characteristics of the processing modules' output (See for example the "3D Device data output control" chapter). Those Processing Module output regions should take the name of the processing module itself followed by their index (such as Scan3dExtractionX).
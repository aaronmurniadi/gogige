|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

ComponentEnable[Region0][Intensity] = False;

// Setup 3D Processing Module source and parameters.

Scan3dExtractionSelector = Scan3dExtraction0;
Scan3dExtractionSource[Scan3dExtraction0] = Region0;
Scan3dExtractionMethod[Scan3dExtraction0] = Default;

// Scan3dExtraction0 output Region (processed data output size).

RegionSelector = Scan3dExtraction0; // First 3D Extraction output Region.
Width[Scan3dExtraction0] = 800; // 3D Extraction output "frame Width".
Height[Scan3dExtraction0] = 640; // 3D Extraction output "frame Height".
RegionMode[Scan3dExtraction0] = On;

// Enable 3D Range output only.

ComponentSelector = Range;
ComponentEnable[Scan3dExtraction0][Range] = True;
ComponentSelector = Reflectance;
ComponentEnable[Scan3dExtraction0][Reflectance] = False;

// Set Region1 position and size.

RegionSelector = Region1; // Second Sensor Region.
OffsetX[Region1] = 80; // Position of Region on sensor.
OffsetY[Region1] = 500; // Position of Region on sensor.
Width[Region1] = 800; // Number of rows of the sensor to read.
Height[Region1] = 300; // Number of lines of the sensor to read.
RegionMode[Region1] = On; // Region 1 features (Width, Height, ...) active.

// Disable Sensor's Region 1 Intensity component output.

ComponentSelector = Intensity;
ComponentEnable[Region1][Intensity] = False; // No Region1 Intensity out.

// Setup 3D Processing Module source and parameters.

Scan3dExtractionSelector = Scan3dExtraction1;
Scan3dExtractionSource[Scan3dExtraction1] = Region1;
Scan3dExtractionMethod[Scan3dExtraction1] = Default;

// Scan3dExtraction1 output Region (processed data outputsize).

RegionSelector = Scan3dExtraction1; // Second 3D Extraction output Region.
Width[Scan3dExtraction1] = 800; // 3D Extraction output "frame Width".
Height[Scan3dExtraction1] = 900; // 3D Extraction output "frame Height".
RegionMode[Scan3dExtraction1] = On;

// Enable 3D Range component output.

ComponentSelector = Range;
ComponentEnable[Scan3dExtraction1][Range] = True;
ComponentSelector = Reflectance;
ComponentEnable[Scan3dExtraction1][Reflectance] = False;

// Start Acquisition.

AquisitionStart;
--
AquisitionStop;
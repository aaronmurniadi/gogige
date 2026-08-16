// 1. DeviceScanType = Linescan3D.
// 3D profiling default mode (3D Range and Reflectance).
// ***
// Dataflow: Region0 -> Scan3dExtraction0 -> Stream0
// ---
// Region0:
//    Defines the size and position of the sensor's Region to process.
// ---
// Scan3dExtraction0:
//    Processing module that extracts a 3D line profile from the 2D data.
//    It's output Region defines the size of the processed buffer to send out.
// ---

// Sensor Region0.
RegionSelector = Region0; // First Sensor Region.
Width[Region0] = 800; // Number of rows of the sensor to read.
Height[Region0] = 500; // Number of lines of the sensor to read.
OffsetX[Region0] = 100; // Position of Region0 on sensor.
OffsetY[Region0] = 400; // Position of Region0 on sensor.
RegionMode[Region0] = On; // Region 0 features (Width, Height, ...) active.

// Disable Region0 Intensity output.
ComponentSelector = Intensity;
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

// Enable Scan3dExtraction0 3D output (Range and Reflectance).
ComponentSelector = Range;
ComponentEnable[Scan3dExtraction0][Range] = True;
ComponentSelector = Reflectance;
ComponentEnable[Scan3dExtraction0][Reflectance] = True;

AquisitionStart;
...
AquisitionStop;
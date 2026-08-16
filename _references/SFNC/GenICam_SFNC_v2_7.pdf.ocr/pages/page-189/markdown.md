|   | FrameEnd FrameActive FrameBurstStart FrameBurstEnd FrameBurstActive LineStart ExposureStart ExposureEnd ExposureActive MultiSlopeExposureLimit1  |
| --- | --- |

Selects the type of trigger to configure.

See Figure 5-1 and Figure 5-3 for details.

Possible values are:

- **AcquisitionStart**: Selects a trigger that starts the Acquisition of one or many frames according to **AcquisitionMode**.
- **AcquisitionEnd**: Selects a trigger that ends the Acquisition of one or many frames according to **AcquisitionMode**.
- **AcquisitionActive**: Selects a trigger that controls the duration of the Acquisition of one or many frames. The Acquisition is activated when the trigger signal becomes active and terminated when it goes back to the inactive state.
- **FrameStart**: Selects a trigger starting the capture of one frame.
- **FrameEnd**: Selects a trigger ending the capture of one frame (mainly used in linescan mode).
- **FrameActive**: Selects a trigger controlling the duration of one frame (mainly used in linescan mode).
- **FrameBurstStart**: Selects a trigger starting the capture of the bursts of frames in an acquisition. **AcquisitionBurstFrameCount** controls the length of each burst unless a FrameBurstEnd trigger is active. The total number of frames captured is also conditioned by AcquisitionFrameCount if AcquisitionMode is MultiFrame.
- **FrameBurstEnd**: Selects a trigger ending the capture of the bursts of frames in an acquisition.
- **FrameBurstActive**: Selects a trigger controlling the duration of the capture of the bursts of frames in an acquisition.
- **LineStart**: Selects a trigger starting the capture of one Line of a Frame (mainly used in linescan mode).
- **ExposureStart**: Selects a trigger controlling the start of the exposure of one Frame (or Line).
- **ExposureEnd**: Selects a trigger controlling the end of the exposure of one Frame (or Line).
- **ExposureActive**: Selects a trigger controlling the duration of the exposure of one frame (or Line).
- **MultiSlopeExposureLimit1**: Selects a trigger controlling the first duration of a multi-slope exposure. Exposure is continued according to the pre-defined multi-slope settings.
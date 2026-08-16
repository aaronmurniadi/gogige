## 5.2 Acquisition features usage model

The AcquisitionMode controls the mode of acquisition for the device. This mainly affects the number of frames captured in the Acquisition (SingleFrame, MultiFrame or Continuous).

The optional AcquisitionArm command is used to verify and freeze all parameters relevant for the image data capture. It prepares the device for the AcquisitionStart.

The AcquisitionStart command is used to start the Acquisition.

The AcquisitionStop command will stop the Acquisition at the end of the current Frame. It can be used in any acquisition mode and if the camera is waiting for a trigger, the pending Frame will be cancelled.

The AcquisitionAbort command can be used to abort an Acquisition at any time. This will end the capture immediately without completing the current Frame.

AcquisitionFrameCount controls the number of frames that will be captured when AcquisitionMode is MultiFrame.

AcquisitionBurstFrameCount determines the length of each burst to capture if the FrameBurstStart trigger is enabled and the FrameBurstEnd trigger is disabled.

AcquisitionFrameRate controls the rate at which the Frames are captured when TriggerMode is Off.

AcquisitionLineRate controls the rate at which the Lines in each Frame are captured when TriggerMode is Off. This is generally useful for linescan cameras.

AcquisitionStatusSelector and AcquisitionStatus can be used to read the status of the internal acquisition signals. The standard acquisition signals Status are: AcquisitionTriggerWait, AcquisitionActive, AcquisitionTransfer, FrameTriggerWait, FrameActive, ExposureActive (See Figure 5-1 and Figure 5-3).
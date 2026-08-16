|   | Line1AnyEdge, ... LinkTrigger0 (If 0 based), LinkTrigger1, ... LinkSpeedChange ActionLate Error Test Device-specific - GigE Vision Specific: PrimaryApplicationSwitch  |
| --- | --- |

Selects which Event to signal to the host application.

See Figure 5-1 to Figure 5-4 and Figure 19-1 for details..

Possible values are:

- AcquisitionTrigger: Device just received a trigger for the Acquisition of one or many Frames.
- AcquisitionTriggerMissed: Device just missed a trigger for the Acquisition of one or many Frames.
- AcquisitionStart: Device just started the Acquisition of one or many Frames.
- AcquisitionEnd: Device just completed the Acquisition of one or many Frames.
- AcquisitionTransferStart: Device just started the transfer of one or many Frames.
- AcquisitionTransferEnd: Device just completed the transfer of one or many Frames.
- AcquisitionError: Device just detected an error during the active Acquisition.
- FrameBurstStart: Device just started the capture of a burst of Frames.
- FrameBurstEnd: Device just completed the capture of a burst of Frames.
- FrameTrigger: Device just received a trigger to start the capture of one Frame.
- FrameTriggerMissed: Device just missed a trigger to start the capture of one Frame.
- FrameStart: Device just started the capture of one Frame.
- FrameEnd: Device just completed the capture of one Frame.
- FrameTransferStart: Device just started the transfer of one Frame.
- FrameTransferEnd: Device just completed the transfer of one Frame.
- LineTrigger: Device just received a trigger to start the capture of one Line.
- LineTriggerMissed: Device just missed a trigger to start the capture of one Line.
- LineStart: Device just started the capture of one Line.
- LineEnd: Device just completed the capture of one Line.
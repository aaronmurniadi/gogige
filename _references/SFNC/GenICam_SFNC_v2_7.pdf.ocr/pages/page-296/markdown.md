## 15 Event Control

This chapter describes how to control the generation of Events to the host application. An Event is a message that is sent to the host application to notify it of the occurrence of an internal event.

Events are typically used to synchronize the host application with some Events happening in the device. A typical use in machine vision is a host application that waits to be notified of the sensor's exposure end to move the inspected part on a conveyer belt.

**EventSelector** selects which particular Event to control. There are many sources of events such as Acquisition, Timer, Counter and I/O lines.

The standard Acquisition related Events are: **AcquisitionTrigger**, **AcquisitionStart**, **AcquisitionEnd**, **AcquisitionTransferStart**, **AcquisitionTransferEnd**, **AcquisitionError**, **FrameTrigger**, **FrameStart**, **FrameEnd**, **FrameBurstStart**, **FrameBurstEnd**, **FrameTransferStart**, **FrameTransferEnd**, **ExposureStart**, **ExposureEnd** (Figure 5-1 to Figure 5-15).

The standard Counters and Timers related Events are: **Counter0Start**, **Counter0End**, **Counter1Start**, **Counter1End**, ... **Timer0Start**, **Timer0End**, **Timer1Start**, **Timer1End**,...

The standard I/O line Events are: **Line0RisingEdge**, **Line0FallingEdge**, **Line0AnyEdge**, **Line1RisingEdge**, **Line1FallingEdge**, ... Note that the event signal is monitored at the same place as **LineStatus** in the I/O control block (See Figure 9-1). This means that event is checked against the condition after the input inverter.

**EventNotification** is used to enable or disable the notification of the occurrence of the internal event selected by **EventSelector**. If **EventNotification** is **Off**, no event of the selected type is generated.

For each of the events listed in the **EventSelector** enumeration, there must be a corresponding event identifier feature with a standard name (ex: **EventExposureEnd**). The controlling application can rely on this feature to register a callback function to be notified that the event happened. This integer event feature must return the unique identifier value identifying the event on the transport layer.

Also for each Event, there should be, in the **EventControl** category, a sub category grouping all the data members related to the particular event (Ex: **EventExposureEndData**).

The other data members in that category should also follow the naming convention described below (Ex: **EventExposureEndTimestamp**).

The recommended optional data members are:

- Timestamp: Unique timestamp of the Event.
- FrameID: Unique ID of the Frame (or image) that generated the Event (if applicable).
- Followed by any other data related to this particular event.

Therefore, the naming convention for the Event related features is:

For each Event (Ex: **ExposureEnd**):
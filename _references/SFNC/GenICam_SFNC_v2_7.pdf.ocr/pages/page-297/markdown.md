|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

- You should provide, in the EventControl category, an ICategory named:

Event prefix + "EventName" + Data postfix (Ex: EventExposureEndData)

- You must provide an IInteger Event feature that will be used as a unique identifier of the event to register the callback and that is named:

Event prefix +"EventName" (Ex: EventExposureEnd).

- You should provide for each optional data member a corresponding feature named:

Event prefix + "EventName"+"DataMember" (Ex: EventExposureEndTimestamp).

For the ExposureEnd event that would be member of EventSelector, this would give:

ICategory EventExposureEndData

IInteger EventExposureEnd

IInteger EventExposureEndTimeStamp

IInteger EventExposureEndFrameID

...

With the above naming convention, for each Event listed in EventSelector:

- A user always knows the name of the Feature to use to register a call back on that Event.
- The user can take the parent of this feature to find the corresponding Event category.
- In this Event category, the user will find all the features related to this Event.

For example, to do a continuous acquisition and be notified at the end of the exposure period of each frame to move the part and also get the timestamp, the following pseudo-code can be used:

Register(Camera.EventExposureEnd, CallbackDataObject, CallbackFunctionPtr)

Camera.EventSelector = ExposureEnd;

Camera.EventNotification = On;

Camera.AcquisitionMode = Continuous;

Camera.AcquisitionStart();

--

// In the callback of the ExposureEnd event, gets the event timestamp:

Timestamp = Camera.EventExposureEndTimestamp;

--

Camera.AcquisitionStop();
|  GEN<i>CAM |   | emva  |
| --- | --- | --- |
|  Version 2.7.1 | Standard Features Naming Convention  |   |

The matching between sequencer sets and currently taken frames can be realized with chunk mode or events.

## 17.2 Sequencer usage examples

### Example 1: Simple features change while acquiring images

In this example, Set 0 and Set 1 are the main sets for the device (here a camera) to capture images. The only parameters within the sequencer sets are Exposure Time, Gain, Width and the (recommended)

SequencerPathSelector and therefore SequencerSetNext, SequencerTriggerSource.

Set 0 is like a free running mode. After every image capture the camera stays in set 0. If a change on Line0 occurs, the sequencer switches to set 1 but only for 1 image capture. The parameters of the singles sets are the following:

- Set 0:
  - ExposureTime = 4000
  - Gain = 1.0
  - Width = 1024
  - SequencerSetNext[0] = 0
  - SequencerTriggerSource[0] = ExposureEnd
  - SequencerSetNext[1] = 1
  - SequencerTriggerSource[1] = Line0

- Set 1:

  - ExposureTime = 2000
  - Gain = 2.0
  - Width = 512
  - SequencerSetNext[0] = 0
  - SequencerTriggerSource[0] = ExposureEnd

The working diagram is shown in the following figure:
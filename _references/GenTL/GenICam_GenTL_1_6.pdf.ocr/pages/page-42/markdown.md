|  GEN<I>CAM |   | emva  |
| --- | --- | --- |
|  Version 1.6 | GenTL Standard  |   |

## 5 Acquisition Engine

### 5.1 Overview

The acquisition engine is the core of the GenTL data stream. Its task is the transportation itself, which mainly consists of the buffer management.

As stated before, the goal for the acquisition engine is to abstract the underlying data transfer mechanism so that it can be used, if not for all, then for most technologies on the market. The target is to acquire data coming from an input stream into memory buffers provided by the GenTL Consumer or made accessible to the GenTL Consumer. The internal design is up to the individual implementation, but there are a few directives it has to follow.

As an essential management element a GenTL acquisition engine holds a number of internal logical buffer pools.

#### 5.1.1 Announced Buffer Pool

All announced buffers are referenced here and are thus known to the acquisition engine. A buffer is known from the point when it is announced until it is revoked (removed from the acquisition engine). It depends on the GenTL Producer if a buffer may be announced during an ongoing acquisition (see 5.2.2). A buffer will stay in this pool even when it is referenced from other queues/pools like the Input Buffer Pool (see 5.1.2) or the Output Buffer Queue (see 5.1.3) or when it is delivered to the GenTL Consumer until it is revoked.

The order of the buffers in the pool is not defined. The maximum possible number of buffers in this pool is only limited due to system resources. The minimum number of buffers in the pool is one or more depending on the technology or the implementation to allow streaming.

#### 5.1.2 Input Buffer Pool

When the acquisition engine receives data from a device it will fill a buffer from this pool. While a buffer is filled it is removed from the pool and if successfully filled, it is put into the output buffer queue. If the transfer was not successful or if the acquisition has been stopped with ACQ STOP FLAGS KILL specified the buffer is placed into the output buffer queue by default. It is up to the implementor to provide additional buffer handling modes which would hand that partially filled buffer differently.

The order of the buffers in the pool is not defined. Only buffers present in the Announced Buffer Pool can be in this pool. The maximum number of buffers in this pool is the number of announced buffers.

#### 5.1.3 Output Buffer Queue

Once a buffer has been successfully filled, it is placed into this queue. As soon as there is at least one buffer in the output buffer queue a previous registered event object gets signaled and the GenTL Consumer can retrieve the event data and thus can identify the filled buffer.

When the event data is retrieved the associated buffer is removed from the output buffer queue. This also means that the data and thus the buffer can only be retrieved once. After the
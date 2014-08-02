#ifndef __RDP_DEF_H__
#define __RDP_DEF_H__

#if !defined(WIN32) && !defined(WIN64)
#include <sys/types.h>
#else
#ifdef __MINGW__
#include <stdint.h>
#endif
#include <windows.h>
#endif

////////////////////////////////////////////////////////////////////////////////

//if compiling on VC6.0 or pre-WindowsXP systems
//use -DLEGACY_WIN32

//if compiling with MinGW, it only works on XP or above
//use -D_WIN32_WINNT=0x0501


#if defined WIN32 || defined WIN64
#ifndef __MINGW__
// Explicitly define 32-bit and 64-bit numbers
typedef __int32 int32_t;
typedef __int64 int64_t;
typedef unsigned __int32 uint32_t;
#ifndef LEGACY_WIN32
typedef unsigned __int64 uint64_t;
#else
// VC 6.0 does not support unsigned __int64: may cause potential problems.
typedef __int64 uint64_t;
#endif

#ifdef RDP_EXPORTS
#define RDP_API __declspec(dllexport)
#else
#define RDP_API
#endif
#else
#define RDP_API
#endif
#else
#define RDP_API __attribute__ ((visibility("default")))
#endif

 
typedef int RDPSOCKET;

typedef enum RDPEPOLLOPT {
    // this values are defined same as linux epoll.h
    // so that if system values are used by mistake, they should have the same effect
    RDP_EPOLL_IN = 0x1,
    RDP_EPOLL_OUT = 0x4,
    RDP_EPOLL_ERR = 0x8
} RDPEPOLLOPT;

typedef enum RDPSTATUS {
    INIT = 1,
    OPENED,
    LISTENING,
    CONNECTING,
    CONNECTED,
    BROKEN,
    CLOSING,
    CLOSED,
    NONEXIST
} RDPSTATUS;

typedef enum RDPSOCKOPT {
    RDP_MSS,             // the Maximum Transfer Unit
    RDP_SNDSYN,          // if sending is blocking
    RDP_RCVSYN,          // if receiving is blocking
    RDP_CC,              // custom congestion control algorithm
    RDP_FC,		         // Flight flag size (window size)
    RDP_SNDBUF,          // maximum buffer in sending queue
    RDP_RCVBUF,          // RDP receiving buffer size
    RDP_LINGER,          // waiting for unsent data when closing
    UDP_SNDBUF,          // UDP sending buffer size
    UDP_RCVBUF,          // UDP receiving buffer size
    RDP_MAXMSG,          // maximum datagram message size
    RDP_MSGTTL,          // time-to-live of a datagram message
    RDP_RENDEZVOUS,      // rendezvous connection mode
    RDP_SNDTIMEO,        // send() timeout
    RDP_RCVTIMEO,        // recv() timeout
    RDP_REUSEADDR,	     // reuse an existing port or create a new one
    RDP_MAXBW,		     // maximum bandwidth (bytes per second) that the connection can use
    RDP_STATE,		     // current socket state, see RDPSTATUS, read only
    RDP_EVENT,		     // current avalable events associated with the socket
    RDP_SNDDATA,	  	 // size of data in the sending buffer
    RDP_RCVDATA		     // size of data available for recv
} RDPSOCKOPT;

typedef struct RDPPerfMon {
    // global measurements
    int64_t msTimeStamp;                 // time since the RDP entity is started, in milliseconds
    int64_t pktSentTotal;                // total number of sent data packets, including retransmissions
    int64_t pktRecvTotal;                // total number of received packets
    int pktSndLossTotal;                 // total number of lost packets (sender side)
    int pktRcvLossTotal;                 // total number of lost packets (receiver side)
    int pktRetransTotal;                 // total number of retransmitted packets
    int pktSentACKTotal;                 // total number of sent ACK packets
    int pktRecvACKTotal;                 // total number of received ACK packets
    int pktSentNAKTotal;                 // total number of sent NAK packets
    int pktRecvNAKTotal;                 // total number of received NAK packets
    int64_t usSndDurationTotal;	      	 // total time duration when RDP is sending data (idle time exclusive)

    // local measurements
    int64_t pktSent;                     // number of sent data packets, including retransmissions
    int64_t pktRecv;                     // number of received packets
    int pktSndLoss;                      // number of lost packets (sender side)
    int pktRcvLoss;                      // number of lost packets (receiver side)
    int pktRetrans;                      // number of retransmitted packets
    int pktSentACK;                      // number of sent ACK packets
    int pktRecvACK;                      // number of received ACK packets
    int pktSentNAK;                      // number of sent NAK packets
    int pktRecvNAK;                      // number of received NAK packets
    double mbpsSendRate;                 // sending rate in Mb/s
    double mbpsRecvRate;                 // receiving rate in Mb/s
    int64_t usSndDuration;		         // busy sending time (i.e., idle time exclusive)

    // instant measurements
    double usPktSndPeriod;               // packet sending period, in microseconds
    int pktFlowWindow;                   // flow window size, in number of packets
    int pktCongestionWindow;             // congestion window size, in number of packets
    int pktFlightSize;                   // number of packets on flight
    double msRTT;                        // RTT, in milliseconds
    double mbpsBandwidth;                // estimated bandwidth, in Mb/s
    int byteAvailSndBuf;                 // available RDP sender buffer size
    int byteAvailRcvBuf;                 // available RDP receiver buffer size
}RDPPerfMon;
 
#endif


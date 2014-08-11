#ifndef RDP_H
#define RDP_H

#include "lint.h"
#include "rdp_def.h"
#include "platform.h"


#if defined PLATFORM_OS_WINDOWS
#ifndef __MINGW__
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
 
#ifdef __cplusplus
extern "C"
{
#endif
RDP_API i32 rdp_startup(rdp_startup_param* param);
RDP_API i32 rdp_startup_get_param(rdp_startup_param* param);
RDP_API i32 rdp_cleanup();
RDP_API i32 rdp_getsyserror();
RDP_API i32 rdp_getsyserrordesc(i32 err, char* desc, ui32* desc_len);

RDP_API i32 rdp_socket_create(rdp_socket_create_param* param, RDPSOCKET* sock);
RDP_API i32 rdp_socket_get_create_param(RDPSOCKET sock, rdp_socket_create_param* param);
RDP_API i32 rdp_socket_get_state(RDPSOCKET sock, ui32* state);
RDP_API i32 rdp_socket_close(RDPSOCKET sock);
RDP_API i32 rdp_socket_bind(RDPSOCKET sock, const char* ip = 0, ui32 port = 0);
RDP_API i32 rdp_socket_listen(RDPSOCKET sock);
RDP_API i32 rdp_socket_connect(RDPSOCKET sock, const char* ip, ui32 port, ui32 timeout, const ui8* buf, ui16 buf_len, RDPSESSIONID* session_id);

RDP_API i32 rdp_session_close(RDPSOCKET sock, RDPSESSIONID session_id, i32 reason);
RDP_API i32 rdp_session_get_state(RDPSOCKET sock, RDPSESSIONID session_id, ui32* state);
RDP_API i32 rdp_session_send(RDPSOCKET sock, RDPSESSIONID session_id, const ui8* buf, ui16 buf_len, ui32 flags);
RDP_API bool rdp_session_is_in_come(RDPSESSIONID session_id);

RDP_API i32 rdp_udp_send(RDPSOCKET sock, const char* ip, ui32 port, const ui8* buf, ui16 buf_len);

RDP_API i32 rdp_addr_to(const sockaddr* addr, ui32 addrlen, char* ip, ui32* iplen, ui32* port);
#ifdef __cplusplus
}
#endif
#endif
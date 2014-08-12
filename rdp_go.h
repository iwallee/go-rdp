#ifndef RDPGO_H
#define RDPGO_H


#include "lint.h"

typedef ui8 bool;
typedef struct sockaddr {} sockaddr;

#include "rdp_def.h"

void __cdecl __on_connect(  rdp_on_connect_param* param);
bool __cdecl __on_before_accept( rdp_on_before_accept_param* param);
void __cdecl __on_accept( rdp_on_accept_param* param);
void __cdecl __on_disconnect( rdp_on_disconnect_param* param);
void __cdecl __on_recv( rdp_on_recv_param* param);
void __cdecl __on_send( rdp_on_send_param* param);
void __cdecl __on_udp_recv( rdp_on_udp_recv_param* param);
ui32 __cdecl __on_hash_addr( sockaddr* addr, ui32 addrlen);

i32 socket_create(void* fn, rdp_socket_create_param* param, RDPSOCKET* sock);


#endif
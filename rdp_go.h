#ifndef RDPGO_H
#define RDPGO_H

typedef char bool;
#include "lint.h"
#include "rdp_def.h"

void __cdecl on_connect(const rdp_on_connect_param* param)
{
}

bool __cdecl on_before_accept(const rdp_on_before_accept_param* param)
{
}

void __cdecl on_accept(const rdp_on_accept_param* param)
{
}

void __cdecl on_disconnect(const rdp_on_disconnect_param* param)
{
}

void __cdecl on_recv(const rdp_on_recv_param* param)
{
}

void __cdecl on_send(const rdp_on_send_param* param)
{
}

void __cdecl on_udp_recv(const rdp_on_udp_recv_param* param)
{
}

ui32 __cdecl on_hash_addr(const sockaddr* addr, ui32 addrlen)
{
}

#endif
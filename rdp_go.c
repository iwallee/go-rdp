#include "rdp_go.h"

void __cdecl __on_connect(  rdp_on_connect_param* param)
{
    on_connect(param);
}

bool __cdecl __on_before_accept( rdp_on_before_accept_param* param)
{
    return on_before_accept(param);
}

void __cdecl __on_accept( rdp_on_accept_param* param)
{
    on_accept(param);
}

void __cdecl __on_disconnect( rdp_on_disconnect_param* param)
{
    on_disconnect(param);
}

void __cdecl __on_recv( rdp_on_recv_param* param)
{
    on_recv(param);
}

void __cdecl __on_send( rdp_on_send_param* param)
{
    on_send(param);
}

void __cdecl __on_udp_recv( rdp_on_udp_recv_param* param)
{
    on_udp_recv(param);
}

ui32 __cdecl __on_hash_addr( sockaddr* addr, ui32 addrlen)
{
    return on_hash_addr(addr, addrlen);
}

i32 socket_create(void* fn, rdp_socket_create_param* param, RDPSOCKET* sock)
{
    typedef i32 (*__fn)(rdp_socket_create_param* , RDPSOCKET* );
    __fn p = (__fn)fn;
    param->on_connect = __on_connect;
    param->on_before_accept = __on_before_accept;
    param->on_accept = __on_accept;
    param->on_disconnect = __on_disconnect;
    param->on_recv = __on_recv;
    param->on_send = __on_send;
    param->on_udp_recv = __on_udp_recv;
    p(param, sock);
}


#ifndef RDPDEF_H
#define RDPDEF_H

#include "lint.h"

#define RDP_SDK_VERSION 0x00010001
#define RDP_VERSION "0.1.0.1"

//���� (in_come)  :�����ⲿ��������(��������ɫ)
//���� (out_come) :���������ⲿ(�ͻ��˽�ɫ)
//rdp socket ����ͬʱ֧��in_come �� out_come,��ͬʱ���������Ϳͻ���

//->rdp_startup
//->rdp_socket_create
//->rdp_socket_bind
//->[rdp_socket_listen:������ܴ���,��Ҫ���ô˷���]
//   |->�Ự:rdp_socket_connect;rdp_session_send;rdp_session_close;
//   |->�ǻỰrdp_udp_send;
//->rdp_socket_close
//->rdp_cleanup

//rdp socket״̬
typedef enum RDPSOCKETSTATUS {
    RDPSOCKETSTATUS_INIT = 1, // ��ʼ
    RDPSOCKETSTATUS_BINDED,   // �Ѱ�
    RDPSOCKETSTATUS_LISTENING,// ����
} RDPSOCKETSTATUS;

//�Ự״̬
typedef enum RDPSESSIONSTAUS {
    RDPSESSIONSTATUS_INIT = 1,
    RDPSESSIONSTATUS_CONNECTING ,   // �Ự������
    RDPSESSIONSTATUS_CONNECTED,     // �Ự������
} RDPOUTCOMESESSIONSTAUS;

typedef enum RDPERROR {
    RDPERROR_SUCCESS = 0,             //�޴���

    RDPERROR_UNKNOWN = -1,            //δ֪����
    RDPERROR_NOTINIT = -2,            //δ��ʼ�����߳�ʼ��ʧ��
    RDPERROR_INVALIDPARAM =-100,      //��Ч�Ĳ���(��ָ���)
    RDPERROR_SYSERROR,                //ϵͳapi����,��rdp_getsyserror��ȡ������

    RDPERROR_SOCKET_RUNOUT ,           //socket������
    RDPERROR_SOCKET_INVALIDSOCKET ,    //��Ч��socket
    RDPERROR_SOCKET_BADSTATE ,         //�����socket״̬

    RDPERROR_SOCKET_ONCONNECTNOTSET ,   //on_connectδ����
    RDPERROR_SOCKET_ONACCEPTNOTSET ,    //on_acceptδ����
    RDPERROR_SOCKET_ONDISCONNECTNOTSET ,//on_disconnectδ����
    RDPERROR_SOCKET_ONRECVNOTSET,       //on_recvδ����
    RDPERROR_SOCKET_ONUDPRECVNOTSET ,   //on_udp_recvδ����

    RDPERROR_SESSION_INVALIDSESSIONID , //��Ч��sessionid
    RDPERROR_SESSION_BADSTATE ,         //����Ļػ�״̬
    RDPERROR_SESSION_CONNTIMEOUT,       //���ӳ�ʱ
    RDPERROR_SESSION_HEARTBEATTIMEOUT,  //������ʱ
    RDPERROR_SESSION_CONNRESET,         //��������:�Է��ر�socket��
} RDPERROR;

typedef enum RDPSESSIONSENDFLAG{
    RDPSESSIONSENDFLAG_ACK     = 0x01, //ȷ���յ����ݰ�
    RDPSESSIONSENDFLAG_INORDER = 0x10, //��˳���ʹ�
}RDPSESSIONSENDFLAG;

typedef enum RDPSESSIONDISCONNECTRESSON{
    DISCONNECTRESSON_NONE = 0,
}RDPSESSIONDISCONNECTRESSON;

typedef ui32 RDPSOCKET;     // != 0
typedef ui64 RDPSESSIONID;  // != 0
 

struct sockaddr;
typedef struct rdp_on_connect_param{
    void*        userdata;
    i32          err;
    RDPSOCKET    sock;
    RDPSESSIONID session_id;
}rdp_on_connect_param;

typedef struct rdp_on_before_accept_param{
    void*            userdata;
    RDPSOCKET        sock;
    RDPSESSIONID     session_id;
    const sockaddr*  addr;
    ui32             addrlen;
    const ui8*       buf;
    ui32             buf_len;
}rdp_on_before_accept_param;

typedef struct rdp_on_accept_param{
    void*            userdata;
    RDPSOCKET        sock;
    RDPSESSIONID     session_id;
    const sockaddr*  addr;
    ui32             addrlen;
    const ui8*       buf;
    ui32             buf_len;
}rdp_on_accept_param;

typedef struct rdp_on_disconnect_param{
    void*        userdata;
    i32          err;
    ui16         reason;//RDPSESSIONDISCONNECTRESSON
    RDPSOCKET    sock;
    RDPSESSIONID session_id;
}rdp_on_disconnect_param;

typedef struct rdp_on_recv_param{
    void*            userdata;
    RDPSOCKET        sock;
    RDPSESSIONID     session_id;
    const ui8*       buf;
    ui16             buf_len;
}rdp_on_recv_param;

typedef struct rdp_on_send_param{
    void*            userdata;
    i32              err;
    RDPSOCKET        sock;
    RDPSESSIONID     session_id;
    ui32             local_send_queue_size;
    ui32             peer_window_size;
}rdp_on_send_param;

typedef struct rdp_on_udp_recv_param{
    void*            userdata;
    RDPSOCKET        sock;
    const sockaddr*  addr;
    ui32             addrlen;
    const ui8*       buf;
    ui16             buf_len;
}rdp_on_udp_recv_param;


typedef struct rdp_startup_param {
    ui32 version;         // rdp sdk �汾�� RDP_SDK_VERSION
    ui8  max_sock;        // ���rdp socket����(Ӧ��С�ڵ���256),Ĭ��1
    ui16 recv_thread_num; // ���ݽ����߳�����:��̨���ݽ����߳�����,Ĭ��1
    ui32 recv_buf_size;   // ���ݽ��ջ�������С:���ݸ�recvfrom�Ļ�������С,Ĭ��4*1024,��ֵӰ�����ݰ�����ܽ��յĴ�С
    //ip��ַhash����,����Ϊ��
    ui32(__cdecl*on_hash_addr)(const sockaddr* addr, ui32 addrlen);
} rdp_startup_param;

typedef struct rdp_socket_create_param {
    void* userdata;
    bool is_v4;                 // �Ƿ���ipv4
    ui16 ack_timeout;           // ȷ�ϳ�ʱ(�ڴ�ʱ����δ�յ�ȷ�ϰ�,��Ϊ��ʱ,ϵͳ���Զ��ط�),Ĭ��300 ms
    ui16 heart_beat_timeout;    // ������ʱ(ÿ��heart_beat_timeout,���ᷢһ������),Ĭ��180s
    ui16 max_send_queue_size;   // �ѷ��͵���δȷ�ϵĶ��д�С,Ĭ��1024,Ϊ0������,�����Ϊ0,���ﵽmax_send_queue_size��,rdp_session_send��������
    ui16 max_recv_queue_size;   // ���ն�������С,Ĭ��1024,Ϊ0������,�����Ϊ0,���ն���(�������ݰ������Ⱥ�˳������)�����ݰ������ﵽ��ֵʱ,���ܾ����������ݰ�
    ui16 in_session_hash_size;  // ����Ựhash��С,Ĭ��1

    //on_connect �������ӻص�,��������ô˻ص�,����������
    void(__cdecl*on_connect)(const rdp_on_connect_param* param);
    //on_before_accept ���ܴ�������ǰ,����ô˻ص�,����������������,����Ϊ��;����false���ܾ�����������(������Ӧ�����)
    bool(__cdecl*on_before_accept)(const rdp_on_before_accept_param* param);
    //on_accept �������ӻص�,��������ô˻ص�,����������
    void(__cdecl*on_accept)(const rdp_on_accept_param* param);
    //on_disconnect ���ӶϿ��ص�,��������
    void(__cdecl*on_disconnect)(const rdp_on_disconnect_param* param);
    //on_recv ���ݽ��ջص�,��������
    void(__cdecl*on_recv)(const rdp_on_recv_param* param);
    //on_send ����,���ɿ������ݽ��ջص�,���������ݵķ���ʹ��rdp_session_send
    void(__cdecl*on_send)(const rdp_on_send_param* param);
    //on_udp_recv ������,���ɿ������ݽ��ջص�,���������ݵķ���ʹ��rdp_udp_send
    void(__cdecl*on_udp_recv)(const rdp_on_udp_recv_param* param);
} rdp_socket_create_param;

#endif
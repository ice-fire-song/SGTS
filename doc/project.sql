drop table public.t_favour;

drop table t_favour_dir;

drop table public.t_goods;

drop table public.t_goods_img;

drop table public.t_goods_label;

drop table public.t_goods_type;

drop table public.t_private_letter;

drop table public.t_user;

/*==============================================================*/
/* User: public                                                 */
/*==============================================================*/
/*==============================================================*/
/* Table: t_favour                                              */
/*==============================================================*/
create table public.t_favour (
   fid                  SERIAL not null,
   gid                  int4                 not null,
   fd_id                int4                 null,
   uid                  int4                 not null,
   create_time          TIMESTAMP WITH TIME ZONE null default CURRENT_TIMESTAMP,
   constraint PK_T_FAVOUR primary key (fid)
);

comment on table public.t_favour is
'收藏表';

comment on column t_favour.gid is
'货品id';

comment on column t_favour.uid is
'用户id';

comment on column t_favour.create_time is
'收藏时间';

/*==============================================================*/
/* Table: t_favour_dir                                          */
/*==============================================================*/
create table t_favour_dir (
   fd_id                serial               not null,
   foldername           varchar(200)         not null,
   uid                  int4                 not null,
   sketch               varchar(200)                null,
   authority_level      int4                 not null   default '1',
   create_time          TIMESTAMP WITH TIME ZONE null default CURRENT_TIMESTAMP,
   constraint PK_T_FAVOUR_DIR primary key (fd_id)
);

/*==============================================================*/
/* Table: t_goods                                               */
/*==============================================================*/
create table public.t_goods (
   gid                  SERIAL not null,
   uid                  int4                 not null,
   gname                varchar(200)         not null,
   gprice               float                null,
   gdetail              text                 null,
   category_id          int4                 null,
   click_number         int4                 not null default '0',
   status               smallint             not null default '1',
   mobilephone_number   varchar(15)          null,
   gliaison             varchar(50)          null,
   openid               varchar(20)          null,
   qq                   varchar(20)          null,
   release_time         TIMESTAMP WITH TIME ZONE null default CURRENT_TIMESTAMP,
   constraint PK_T_GOODS primary key (gid)
);

comment on table public.t_goods is
'货品表';

comment on column t_goods.gid is
'货品id';

comment on column t_goods.uid is
'用户id';

comment on column t_goods.gname is
'货品名称';

comment on column t_goods.gprice is
'货品价格';

comment on column t_goods.gdetail is
'货品详情';

comment on column t_goods.category_id is
'类别id, 1 商品、2 免费商品、3 需求';

comment on column t_goods.click_number is
'点击量';

comment on column t_goods.status is
'货品状态（0 删除、1 上架、2下架）';

comment on column t_goods.mobilephone_number is
'手机号码';

comment on column t_goods.gliaison is
'联系人';

comment on column t_goods.openid is
'微信号';

comment on column t_goods.qq is
'QQ号';

comment on column t_goods.release_time is
'发布时间';

/*==============================================================*/
/* Table: t_goods_img                                           */
/*==============================================================*/
create table public.t_goods_img (
   id                   SERIAL not null,
   gid                  int4                 not null,
   image_name           varchar(100)         null,
   image_ext            varchar(10)          null,
   save_path            varchar(255)         not null,
   image_size           float                null,
   release_time         TIMESTAMP WITH TIME ZONE null default CURRENT_TIMESTAMP,
   constraint PK_T_GOODS_IMG primary key (id)
);

comment on table public.t_goods_img is
'货品图片表';

comment on column t_goods_img.gid is
'货品id';

comment on column t_goods_img.image_name is
'图片名称';

comment on column t_goods_img.image_ext is
'图片扩展名';

comment on column t_goods_img.save_path is
'图片保存路径';

comment on column t_goods_img.image_size is
'图片大小';

comment on column t_goods_img.release_time is
'发布时间';

/*==============================================================*/
/* Table: t_goods_label                                         */
/*==============================================================*/
create table public.t_goods_label (
   gid                  int4                 not null,
   label_name           varchar(150)         not null,
   create_time          TIMESTAMP WITH TIME ZONE null default CURRENT_TIMESTAMP,
   constraint PK_T_LABEL primary key (gid, label_name)
);

comment on table public.t_goods_label is
'货品标签表';

comment on column t_goods_label.gid is
'货品id';

comment on column t_goods_label.label_name is
'标签';

comment on column t_goods_label.create_time is
'创建时间';
/*==============================================================*/
/* Table: t_goods_type                                        */
/*==============================================================*/
create table public.t_goods_type (
   gt_id                  serial                not null,
   type_name              varchar(150)         not null,
   create_time            TIMESTAMP WITH TIME ZONE null default CURRENT_TIMESTAMP,
   constraint PK_T_GOODS_TYPE primary key (gt_id, type_name)
);

comment on table public.t_goods_type is
'货品种类表';

comment on column t_goods_type.gt_id is
'货品种类id';

comment on column t_goods_type.type_name is
'货品种类';

comment on column t_goods_type.create_time is
'创建时间';
/*==============================================================*/
/* Table: t_private_letter                                      */
/*==============================================================*/
create table public.t_private_letter (
   plid                 SERIAL not null,
   user_id              int4                 null,
   friend_id            int4                 null,
   sender_id            int4                 null,
   receiver_id          int4                 null,
   massage_type         smallint             null,
   message_content      text                 null,
   send_time            TIMESTAMP WITH TIME ZONE null default CURRENT_TIMESTAMP,
   status               smallint             not null default '1',
   constraint PK_T_PRIVATE_LETTER primary key (plid)
);

comment on table public.t_private_letter is
'私信表';

comment on column t_private_letter.plid is
'私信id';

comment on column t_private_letter.user_id is
'非真实发送者';

comment on column t_private_letter.friend_id is
'非真实接收者';

comment on column t_private_letter.sender_id is
'发送者';

comment on column t_private_letter.receiver_id is
'接收者';

comment on column t_private_letter.massage_type is
'信息类型';

comment on column t_private_letter.message_content is
'信息内容';

comment on column t_private_letter.send_time is
'发送时间';

comment on column t_private_letter.status is
'消息状态 1：未读 2：已读 3：删除';

/*==============================================================*/
/* Table: t_user                                                */
/*==============================================================*/
create table public.t_user (
   uid                  SERIAL not null,
   username             varchar(100)         not null,
   password             varchar(50)          not null,
   user_role            smallint             not null,
   head_sculpture_path  varchar(255)         not null default '/',
   label                text                 null,
   create_time          TIMESTAMP WITH TIME ZONE null default CURRENT_TIMESTAMP,
   status               smallint             not null default '0',
   constraint PK_T_USER primary key (uid)
);

comment on table public.t_user is
'用户表';

comment on column t_user.uid is
'用户id';

comment on column t_user.username is
'用户名';

comment on column t_user.password is
'密码';

comment on column t_user.user_role is
'用户角色';

comment on column t_user.head_sculpture_path is
'用户头像';

comment on column t_user.label is
'标签';

comment on column t_user.create_time is
'注册时间';

comment on column t_user.status is
'用户状态（0 正常，1 封禁）';

alter table t_favour
   add constraint FK_T_FAVOUR_REFERENCE_T_GOODS foreign key (gid)
      references t_goods (gid)
      on delete restrict on update restrict;

alter table t_favour
   add constraint FK_T_FAVOUR_REFERENCE_T_FAVOUR foreign key (fd_id)
      references t_favour_dir (fd_id)
      on delete restrict on update restrict;

alter table t_goods
   add constraint FK_T_GOODS_REFERENCE_T_USER foreign key (uid)
      references t_user (uid)
      on delete restrict on update restrict;

alter table t_goods_img
   add constraint FK_T_GOODS__REFERENCE_T_GOODS foreign key (gid)
      references t_goods (gid)
      on delete restrict on update restrict;

alter table t_goods_label
   add constraint FK_T_LABEL_REFERENCE_T_GOODS foreign key (gid)
      references t_goods (gid)
      on delete restrict on update restrict;

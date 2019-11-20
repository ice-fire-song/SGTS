if exists(select 1 from sys.sysforeignkey where role='FK_T_FAVOUR_REFERENCE_T_GOODS') then
    alter table t_favour
       delete foreign key FK_T_FAVOUR_REFERENCE_T_GOODS
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_T_GOODS_REFERENCE_T_USER') then
    alter table t_goods
       delete foreign key FK_T_GOODS_REFERENCE_T_USER
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_T_GOODS__REFERENCE_T_GOODS') then
    alter table t_goods_img
       delete foreign key FK_T_GOODS__REFERENCE_T_GOODS
end if;

if exists(select 1 from sys.sysforeignkey where role='FK_T_LABEL_REFERENCE_T_GOODS') then
    alter table t_label
       delete foreign key FK_T_LABEL_REFERENCE_T_GOODS
end if;

drop table if exists t_favour;

drop table if exists t_goods;

drop table if exists t_goods_img;

drop table if exists t_label;

drop table if exists t_private_letter;

drop table if exists t_user;

/*==============================================================*/
/* Table: t_favour                                              */
/*==============================================================*/
create table t_favour 
(
   fid                  serial                         not null,
   gid                  serail                         null,
   uid                  int4                           null,
   create_time          timestamp with time zone       null,
   constraint PK_T_FAVOUR primary key (fid)
);

/*==============================================================*/
/* Table: t_goods                                               */
/*==============================================================*/
create table t_goods 
(
   gid                  serail                         not null,
   uid                  int4                           null,
   gname                varchar(200)                   not null,
   gprice               float                          not null,
   gdetail              text                           null,
   mobilephone_number   varchar(15)                    not null,
   category_id          int4                           null,
   click_number         int                            null,
   status               smallint                       null,
   gliaison             varchar(50)                    null,
   openid               varchar(20)                    null,
   qq                   varchar(20)                    null,
   release_time         timestamp with time zone       null,
   constraint PK_T_GOODS primary key (gid)
);

/*==============================================================*/
/* Table: t_goods_img                                           */
/*==============================================================*/
create table t_goods_img 
(
   id                   serial                         not null,
   gid                  serail                         null,
   image_name           varchar(100)                   null,
   image_ext            varchar(10)                    null,
   save_path            varchar(255)                   null,
   image_size           float                          null,
   release_time         timestamp with time zone       null,
   constraint PK_T_GOODS_IMG primary key (id)
);

/*==============================================================*/
/* Table: t_label                                               */
/*==============================================================*/
create table t_label 
(
   id                   serial                         null,
   gid                  int4                           not null,
   label_name           varchar(150)                   not null,
   t_g_gid              serail                         null,
   create_time          timestamp with time zone       null,
   constraint PK_T_LABEL primary key (gid, label_name)
);

/*==============================================================*/
/* Table: t_private_letter                                      */
/*==============================================================*/
create table t_private_letter 
(
   plid                 serial                         not null,
   user_id              int4                           null,
   friend_id            int4                           null,
   sender_id            int4                           null,
   receiver_id          int4                           null,
   massage_type         smallint                       null,
   message_content      text                           null,
   send_time            timestamp with time zone       null,
   status               smallint                       null,
   constraint PK_T_PRIVATE_LETTER primary key(plid)
);

/*==============================================================*/
/* Table: t_user                                                */
/*==============================================================*/
create table t_user 
(
   uid                  serial                         not null,
   username             varchar(100)                   not null,
   password             varchar(50)                    null,
   head_sculpture_path  varchar(255)                   null,
   label                text                           null,
   create_time          timestamp with time zone       null,
   constraint PK_T_USER primary key clustered (uid, username)
);

comment on table t_user is 
'用户表';

alter table t_favour
   add constraint FK_T_FAVOUR_REFERENCE_T_GOODS foreign key (gid)
      references t_goods (gid)
      on update restrict
      on delete restrict;

alter table t_goods
   add constraint FK_T_GOODS_REFERENCE_T_USER foreign key (uid, )
      references t_user (uid, username)
      on update restrict
      on delete restrict;

alter table t_goods_img
   add constraint FK_T_GOODS__REFERENCE_T_GOODS foreign key (gid)
      references t_goods (gid)
      on update restrict
      on delete restrict;

alter table t_label
   add constraint FK_T_LABEL_REFERENCE_T_GOODS foreign key (t_g_gid)
      references t_goods (gid)
      on update restrict
      on delete restrict;

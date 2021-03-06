        z            Market    14.1    14.1 /               0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                        0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            !           1262    24614    Market    DATABASE     c   CREATE DATABASE "Market" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'Thai_Thailand.874';
    DROP DATABASE "Market";
                postgres    false            @           1247    24616    option_product    TYPE     E   CREATE TYPE public.option_product AS ENUM (
    'buy',
    'sell'
);
 !   DROP TYPE public.option_product;
       public          postgres    false            Ñ            1259    24621    interested_post    TABLE     d   CREATE TABLE public.interested_post (
    post_id integer NOT NULL,
    user_id integer NOT NULL
);
 #   DROP TABLE public.interested_post;
       public         heap    postgres    false            Ò            1259    24624    interested_post_post_id_seq    SEQUENCE        CREATE SEQUENCE public.interested_post_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 2   DROP SEQUENCE public.interested_post_post_id_seq;
       public          postgres    false    209            "           0    0    interested_post_post_id_seq    SEQUENCE OWNED BY     [   ALTER SEQUENCE public.interested_post_post_id_seq OWNED BY public.interested_post.post_id;
          public          postgres    false    210            Ó            1259    24625    interested_post_user_id_seq    SEQUENCE        CREATE SEQUENCE public.interested_post_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 2   DROP SEQUENCE public.interested_post_user_id_seq;
       public          postgres    false    209            #           0    0    interested_post_user_id_seq    SEQUENCE OWNED BY     [   ALTER SEQUENCE public.interested_post_user_id_seq OWNED BY public.interested_post.user_id;
          public          postgres    false    211            Ô            1259    24626    pin    TABLE     n   CREATE TABLE public.pin (
    user_id integer,
    pin_id integer NOT NULL,
    pin character varying(255)
);
    DROP TABLE public.pin;
       public         heap    postgres    false            Û            1259    24780    pin_pin_id_seq    SEQUENCE        CREATE SEQUENCE public.pin_pin_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.pin_pin_id_seq;
       public          postgres    false    212            $           0    0    pin_pin_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.pin_pin_id_seq OWNED BY public.pin.pin_id;
          public          postgres    false    219            Õ            1259    24630    post    TABLE     ñ   CREATE TABLE public.post (
    post_id integer NOT NULL,
    pin_id integer,
    product_name character varying(255),
    post_date date,
    product_option public.option_product,
    price integer,
    amount integer,
    tag_id integer
);
    DROP TABLE public.post;
       public         heap    postgres    false    832            Ö            1259    24633    post_post_id_seq    SEQUENCE        CREATE SEQUENCE public.post_post_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 '   DROP SEQUENCE public.post_post_id_seq;
       public          postgres    false    213            %           0    0    post_post_id_seq    SEQUENCE OWNED BY     E   ALTER SEQUENCE public.post_post_id_seq OWNED BY public.post.post_id;
          public          postgres    false    214            ×            1259    24635    tag    TABLE     Y   CREATE TABLE public.tag (
    tag_id integer NOT NULL,
    tag_name character varying
);
    DROP TABLE public.tag;
       public         heap    postgres    false            Ø            1259    24640    tag_tag_id_seq    SEQUENCE        CREATE SEQUENCE public.tag_tag_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 %   DROP SEQUENCE public.tag_tag_id_seq;
       public          postgres    false    215            &           0    0    tag_tag_id_seq    SEQUENCE OWNED BY     A   ALTER SEQUENCE public.tag_tag_id_seq OWNED BY public.tag.tag_id;
          public          postgres    false    216            Ù            1259    24641    user1    TABLE     Í   CREATE TABLE public.user1 (
    user_id integer NOT NULL,
    name character varying(255),
    address character varying(255),
    contact character varying(255),
    phone_number character varying(10)
);
    DROP TABLE public.user1;
       public         heap    postgres    false            Ú            1259    24646    user1_user_id_seq    SEQUENCE        CREATE SEQUENCE public.user1_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.user1_user_id_seq;
       public          postgres    false    217            '           0    0    user1_user_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.user1_user_id_seq OWNED BY public.user1.user_id;
          public          postgres    false    218            t           2604    24647    interested_post post_id    DEFAULT        ALTER TABLE ONLY public.interested_post ALTER COLUMN post_id SET DEFAULT nextval('public.interested_post_post_id_seq'::regclass);
 F   ALTER TABLE public.interested_post ALTER COLUMN post_id DROP DEFAULT;
       public          postgres    false    210    209            u           2604    24648    interested_post user_id    DEFAULT        ALTER TABLE ONLY public.interested_post ALTER COLUMN user_id SET DEFAULT nextval('public.interested_post_user_id_seq'::regclass);
 F   ALTER TABLE public.interested_post ALTER COLUMN user_id DROP DEFAULT;
       public          postgres    false    211    209            v           2604    24781 
   pin pin_id    DEFAULT     h   ALTER TABLE ONLY public.pin ALTER COLUMN pin_id SET DEFAULT nextval('public.pin_pin_id_seq'::regclass);
 9   ALTER TABLE public.pin ALTER COLUMN pin_id DROP DEFAULT;
       public          postgres    false    219    212            w           2604    24650    post post_id    DEFAULT     l   ALTER TABLE ONLY public.post ALTER COLUMN post_id SET DEFAULT nextval('public.post_post_id_seq'::regclass);
 ;   ALTER TABLE public.post ALTER COLUMN post_id DROP DEFAULT;
       public          postgres    false    214    213            x           2604    24652 
   tag tag_id    DEFAULT     h   ALTER TABLE ONLY public.tag ALTER COLUMN tag_id SET DEFAULT nextval('public.tag_tag_id_seq'::regclass);
 9   ALTER TABLE public.tag ALTER COLUMN tag_id DROP DEFAULT;
       public          postgres    false    216    215            y           2604    24653    user1 user_id    DEFAULT     n   ALTER TABLE ONLY public.user1 ALTER COLUMN user_id SET DEFAULT nextval('public.user1_user_id_seq'::regclass);
 <   ALTER TABLE public.user1 ALTER COLUMN user_id DROP DEFAULT;
       public          postgres    false    218    217                      0    24621    interested_post 
   TABLE DATA           ;   COPY public.interested_post (post_id, user_id) FROM stdin;
    public          postgres    false    209   +2                 0    24626    pin 
   TABLE DATA           3   COPY public.pin (user_id, pin_id, pin) FROM stdin;
    public          postgres    false    212   H2                 0    24630    post 
   TABLE DATA           o   COPY public.post (post_id, pin_id, product_name, post_date, product_option, price, amount, tag_id) FROM stdin;
    public          postgres    false    213   ±2                 0    24635    tag 
   TABLE DATA           /   COPY public.tag (tag_id, tag_name) FROM stdin;
    public          postgres    false    215   ö2                 0    24641    user1 
   TABLE DATA           N   COPY public.user1 (user_id, name, address, contact, phone_number) FROM stdin;
    public          postgres    false    217   3       (           0    0    interested_post_post_id_seq    SEQUENCE SET     J   SELECT pg_catalog.setval('public.interested_post_post_id_seq', 1, false);
          public          postgres    false    210            )           0    0    interested_post_user_id_seq    SEQUENCE SET     J   SELECT pg_catalog.setval('public.interested_post_user_id_seq', 1, false);
          public          postgres    false    211            *           0    0    pin_pin_id_seq    SEQUENCE SET     <   SELECT pg_catalog.setval('public.pin_pin_id_seq', 2, true);
          public          postgres    false    219            +           0    0    post_post_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.post_post_id_seq', 14, true);
          public          postgres    false    214            ,           0    0    tag_tag_id_seq    SEQUENCE SET     =   SELECT pg_catalog.setval('public.tag_tag_id_seq', 31, true);
          public          postgres    false    216            -           0    0    user1_user_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.user1_user_id_seq', 22, true);
          public          postgres    false    218            {           2606    24783    pin pin_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.pin
    ADD CONSTRAINT pin_pkey PRIMARY KEY (pin_id);
 6   ALTER TABLE ONLY public.pin DROP CONSTRAINT pin_pkey;
       public            postgres    false    212            }           2606    24657    post post_pkey 
   CONSTRAINT     Q   ALTER TABLE ONLY public.post
    ADD CONSTRAINT post_pkey PRIMARY KEY (post_id);
 8   ALTER TABLE ONLY public.post DROP CONSTRAINT post_pkey;
       public            postgres    false    213                       2606    24659    tag tag_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.tag
    ADD CONSTRAINT tag_pkey PRIMARY KEY (tag_id);
 6   ALTER TABLE ONLY public.tag DROP CONSTRAINT tag_pkey;
       public            postgres    false    215                       2606    24661    user1 user1_pkey 
   CONSTRAINT     S   ALTER TABLE ONLY public.user1
    ADD CONSTRAINT user1_pkey PRIMARY KEY (user_id);
 :   ALTER TABLE ONLY public.user1 DROP CONSTRAINT user1_pkey;
       public            postgres    false    217                       2606    24788    post fk_pin_post    FK CONSTRAINT     p   ALTER TABLE ONLY public.post
    ADD CONSTRAINT fk_pin_post FOREIGN KEY (pin_id) REFERENCES public.pin(pin_id);
 :   ALTER TABLE ONLY public.post DROP CONSTRAINT fk_pin_post;
       public          postgres    false    3195    213    212                       2606    24667    pin pin_user_id_fkey    FK CONSTRAINT     x   ALTER TABLE ONLY public.pin
    ADD CONSTRAINT pin_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.user1(user_id);
 >   ALTER TABLE ONLY public.pin DROP CONSTRAINT pin_user_id_fkey;
       public          postgres    false    217    3201    212                       2606    24672    interested_post post_id    FK CONSTRAINT     z   ALTER TABLE ONLY public.interested_post
    ADD CONSTRAINT post_id FOREIGN KEY (post_id) REFERENCES public.post(post_id);
 A   ALTER TABLE ONLY public.interested_post DROP CONSTRAINT post_id;
       public          postgres    false    209    213    3197                       2606    24687    interested_post user_id    FK CONSTRAINT     {   ALTER TABLE ONLY public.interested_post
    ADD CONSTRAINT user_id FOREIGN KEY (user_id) REFERENCES public.user1(user_id);
 A   ALTER TABLE ONLY public.interested_post DROP CONSTRAINT user_id;
       public          postgres    false    3201    209    217                  xÑãââ Å ©         Y   x32äB#3S3c.##N#N£DC0×²Ür_}³¨Ð(Ï¬Ü ÷½ DÀÄbË
3°d§¢Ð\s½Òô´ÔÄâJ®=... kéÒ         5   x34æ4äÌ.­ä4202Ò50Ò54äLrMMM9-8¸M8ð(0äÑãââ ]ö            x36àÌHÌHä26Ð1z\\\ 8Íº         .   x32ä,(J¬,-ÉàLJÌKÏÎÏæL,.æ44242æ22Â#£ÇÅÅ rÉ      
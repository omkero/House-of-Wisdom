PGDMP  3                     }         	   wisdom_db    17.2    17.2     3           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                           false            4           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                           false            5           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                           false            6           1262    16460 	   wisdom_db    DATABASE     �   CREATE DATABASE wisdom_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United Kingdom.1252';
    DROP DATABASE wisdom_db;
                     postgres    false            �            1259    16461    articles    TABLE     1  CREATE TABLE public.articles (
    article_id integer NOT NULL,
    article_title character varying(255) NOT NULL,
    article_content text NOT NULL,
    created_at timestamp without time zone DEFAULT now(),
    banner_img text,
    tags jsonb,
    views integer DEFAULT 0,
    is_hidden boolean DEFAULT false,
    user_id integer,
    username character varying(255),
    updated_at timestamp without time zone DEFAULT now(),
    public_article_title character varying(255) DEFAULT ''::character varying,
    description text DEFAULT 'no description'::text
);
    DROP TABLE public.articles;
       public         heap r       postgres    false            �            1259    16472    articles_article_id_seq    SEQUENCE     �   CREATE SEQUENCE public.articles_article_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 .   DROP SEQUENCE public.articles_article_id_seq;
       public               postgres    false    217            7           0    0    articles_article_id_seq    SEQUENCE OWNED BY     S   ALTER SEQUENCE public.articles_article_id_seq OWNED BY public.articles.article_id;
          public               postgres    false    218            �            1259    16473    users    TABLE     �  CREATE TABLE public.users (
    user_id integer NOT NULL,
    username character varying(255) NOT NULL,
    email_address character varying(255),
    password character varying(255),
    is_banned boolean DEFAULT false,
    is_suspended boolean DEFAULT false,
    articles_count integer DEFAULT 0,
    profile_img text DEFAULT ''::text,
    created_at timestamp without time zone DEFAULT now()
);
    DROP TABLE public.users;
       public         heap r       postgres    false            �            1259    16483    users_user_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.users_user_id_seq;
       public               postgres    false    219            8           0    0    users_user_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;
          public               postgres    false    220            �           2604    16484    articles article_id    DEFAULT     z   ALTER TABLE ONLY public.articles ALTER COLUMN article_id SET DEFAULT nextval('public.articles_article_id_seq'::regclass);
 B   ALTER TABLE public.articles ALTER COLUMN article_id DROP DEFAULT;
       public               postgres    false    218    217            �           2604    16485    users user_id    DEFAULT     n   ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);
 <   ALTER TABLE public.users ALTER COLUMN user_id DROP DEFAULT;
       public               postgres    false    220    219            -          0    16461    articles 
   TABLE DATA           �   COPY public.articles (article_id, article_title, article_content, created_at, banner_img, tags, views, is_hidden, user_id, username, updated_at, public_article_title, description) FROM stdin;
    public               postgres    false    217          /          0    16473    users 
   TABLE DATA           �   COPY public.users (user_id, username, email_address, password, is_banned, is_suspended, articles_count, profile_img, created_at) FROM stdin;
    public               postgres    false    219   7       9           0    0    articles_article_id_seq    SEQUENCE SET     F   SELECT pg_catalog.setval('public.articles_article_id_seq', 66, true);
          public               postgres    false    218            :           0    0    users_user_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.users_user_id_seq', 29, true);
          public               postgres    false    220            �           2606    16487    users email_unique 
   CONSTRAINT     V   ALTER TABLE ONLY public.users
    ADD CONSTRAINT email_unique UNIQUE (email_address);
 <   ALTER TABLE ONLY public.users DROP CONSTRAINT email_unique;
       public                 postgres    false    219            �           2606    16489    users username_unique 
   CONSTRAINT     T   ALTER TABLE ONLY public.users
    ADD CONSTRAINT username_unique UNIQUE (username);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT username_unique;
       public                 postgres    false    219            -      x������ � �      /      x������ � �     
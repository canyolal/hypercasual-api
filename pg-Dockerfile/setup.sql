CREATE DATABASE hypercasual ENCODING = 'UTF8' LC_COLLATE = 'en_US.utf8' LC_CTYPE = 'en_US.utf8';

ALTER DATABASE hypercasual OWNER TO postgres;

\connect hypercasual

CREATE EXTENSION citext;

CREATE TABLE IF NOT EXISTS publishers (
    id bigserial PRIMARY KEY,
    name text UNIQUE NOT NULL,
    link text NOT NULL,
    version integer NOT NULL DEFAULT 1
);

INSERT INTO publishers (name, link) 
VALUES
    ('Matchingham', 'https://apps.apple.com/tr/developer/matchingham-games/id1513009812?see-all=i-phonei-pad-apps'),
    ('Miniclip','https://apps.apple.com/us/developer/miniclip-com/id337457683?see-all=i-phonei-pad-apps'),
    ('Voodoo','https://apps.apple.com/us/developer/voodoo/id714804730?see-all=i-phonei-pad-apps'),
    ('Good Job Games','https://apps.apple.com/tr/developer/good-job-games/id1191495496?see-all=i-phonei-pad-apps'),
    ('Ketchapp','https://apps.apple.com/us/developer/ketchapp/id528065807?see-all=i-phonei-pad-apps'),
    ('Alictus','https://apps.apple.com/tr/developer/alictus/id892399717?see-all=i-phonei-pad-apps'),
    ('Lion Studios','https://apps.apple.com/us/developer/lion-studios/id1362220666?see-all=i-phonei-pad-apps'),
    ('Rollic','https://apps.apple.com/us/developer/rollic-games/id1452111779?see-all=i-phonei-pad-apps'),
    ('Kwalee','https://apps.apple.com/tr/developer/kwalee-ltd/id497961736?see-all=i-phonei-pad-apps'),
    ('BoomBit','https://apps.apple.com/kh/developer/boombit-inc/id1045926022?see-all=i-phonei-pad-apps'),
    ('Amanotes','https://apps.apple.com/us/developer/amanotes-pte-ltd/id1441389613?see-all=i-phonei-pad-apps'),
    ('Azur Games','https://apps.apple.com/us/developer/azur-interactive-games-limited/id1296347323?see-all=i-phonei-pad-apps'),
    ('Crazy Labs','https://apps.apple.com/us/developer/crazy-labs/id721307559?see-all=i-phonei-pad-apps'),
    ('Coda','https://apps.apple.com/us/developer/coda-platform-limited/id1475474579?see-all=i-phonei-pad-apps'),
    ('Ducky','https://apps.apple.com/us/developer/ducky-games/id957096633?see-all=i-phonei-pad-apps'),
    ('Gismart','https://apps.apple.com/us/developer/gismart/id666830030?see-all=i-phonei-pad-apps'),
    ('Green Panda Games','https://apps.apple.com/tr/developer/green-panda-games/id669978473?see-all=i-phonei-pad-apps'),
    ('Homa','https://apps.apple.com/tr/developer/homa-games/id1508492426?see-all=i-phonei-pad-apps'),
    ('JoyPac','https://apps.apple.com/tr/developer/joypac/id1422558565?see-all=i-phonei-pad-apps'),
    ('Moonee','https://apps.apple.com/us/developer/moonee-publishing-ltd/id1469957859?see-all=i-phonei-pad-apps'),
    ('Playgendary','https://apps.apple.com/us/developer/playgendary-limited/id1487320337?see-all=i-phonei-pad-apps'),
    ('SayGames','https://apps.apple.com/tr/developer/saygames-ltd/id1551847165?see-all=i-phonei-pad-apps'),
    ('Supersonic','https://apps.apple.com/us/developer/supersonic-studios-ltd/id1499845738?see-all=i-phonei-pad-apps'),
    ('TapNation','https://apps.apple.com/tr/developer/tapnation/id1483575279?see-all=i-phonei-pad-apps'),
    ('Tastypill','https://apps.apple.com/us/developer/tastypill/id1022434729?see-all=i-phonei-pad-apps'),
    ('Yso Corp','https://apps.apple.com/us/developer/yso-corp/id659815325?see-all=i-phonei-pad-apps'),
    ('Zplay','https://apps.apple.com/tr/developer/zplay-beijing-info-tech-co-ltd/id531022725?see-all=i-phonei-pad-apps');

CREATE TABLE IF NOT EXISTS games (
	id bigserial PRIMARY KEY, 
	name text NOT NULL,
	genre text NOT NULL,
	publisher_name text NOT NULL,
	created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
	version integer NOT NULL DEFAULT 1
);

CREATE TABLE IF NOT EXISTS maillist (
    id bigserial PRIMARY KEY,
    email citext UNIQUE NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
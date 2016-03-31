--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;

--
-- Name: recitation; Type: SCHEMA; Schema: -; Owner: david
--

CREATE SCHEMA recitation;


ALTER SCHEMA recitation OWNER TO david;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = recitation, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: course; Type: TABLE; Schema: recitation; Owner: david; Tablespace: 
--

CREATE TABLE course (
    id integer NOT NULL,
    name text NOT NULL,
    numtracks integer NOT NULL
);


ALTER TABLE course OWNER TO david;

--
-- Name: course_id_seq; Type: SEQUENCE; Schema: recitation; Owner: david
--

CREATE SEQUENCE course_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE course_id_seq OWNER TO david;

--
-- Name: course_id_seq; Type: SEQUENCE OWNED BY; Schema: recitation; Owner: david
--

ALTER SEQUENCE course_id_seq OWNED BY course.id;


--
-- Name: problem; Type: TABLE; Schema: recitation; Owner: david; Tablespace: 
--

CREATE TABLE problem (
    cid integer NOT NULL,
    recitation text NOT NULL,
    problem text NOT NULL,
    compulsory text
);


ALTER TABLE problem OWNER TO david;

--
-- Name: recitation; Type: TABLE; Schema: recitation; Owner: david; Tablespace: 
--

CREATE TABLE recitation (
    cid integer NOT NULL,
    name text NOT NULL
);


ALTER TABLE recitation OWNER TO david;

--
-- Name: solved; Type: TABLE; Schema: recitation; Owner: david; Tablespace: 
--

CREATE TABLE solved (
    sid integer NOT NULL,
    cid integer NOT NULL,
    recitation text NOT NULL,
    problem text NOT NULL,
    letter text NOT NULL,
    called boolean,
    points integer
);


ALTER TABLE solved OWNER TO david;

--
-- Name: student; Type: TABLE; Schema: recitation; Owner: david; Tablespace: 
--

CREATE TABLE student (
    id integer NOT NULL,
    name text,
    password text
);


ALTER TABLE student OWNER TO david;

--
-- Name: student_id_seq; Type: SEQUENCE; Schema: recitation; Owner: david
--

CREATE SEQUENCE student_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE student_id_seq OWNER TO david;

--
-- Name: student_id_seq; Type: SEQUENCE OWNED BY; Schema: recitation; Owner: david
--

ALTER SEQUENCE student_id_seq OWNED BY student.id;


--
-- Name: subproblem; Type: TABLE; Schema: recitation; Owner: david; Tablespace: 
--

CREATE TABLE subproblem (
    cid integer NOT NULL,
    recitation text NOT NULL,
    problem text NOT NULL,
    letter text NOT NULL
);


ALTER TABLE subproblem OWNER TO david;

--
-- Name: takes; Type: TABLE; Schema: recitation; Owner: david; Tablespace: 
--

CREATE TABLE takes (
    cid integer NOT NULL,
    sid integer NOT NULL
);


ALTER TABLE takes OWNER TO david;

--
-- Name: track; Type: TABLE; Schema: recitation; Owner: david; Tablespace: 
--

CREATE TABLE track (
    cid integer NOT NULL,
    name text NOT NULL,
    sid integer NOT NULL,
    track integer
);


ALTER TABLE track OWNER TO david;

--
-- Name: id; Type: DEFAULT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY course ALTER COLUMN id SET DEFAULT nextval('course_id_seq'::regclass);


--
-- Name: id; Type: DEFAULT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY student ALTER COLUMN id SET DEFAULT nextval('student_id_seq'::regclass);


--
-- Name: course_pkey; Type: CONSTRAINT; Schema: recitation; Owner: david; Tablespace: 
--

ALTER TABLE ONLY course
    ADD CONSTRAINT course_pkey PRIMARY KEY (id);


--
-- Name: problem_pkey; Type: CONSTRAINT; Schema: recitation; Owner: david; Tablespace: 
--

ALTER TABLE ONLY problem
    ADD CONSTRAINT problem_pkey PRIMARY KEY (cid, recitation, problem);


--
-- Name: recitation_pkey; Type: CONSTRAINT; Schema: recitation; Owner: david; Tablespace: 
--

ALTER TABLE ONLY recitation
    ADD CONSTRAINT recitation_pkey PRIMARY KEY (cid, name);


--
-- Name: solved_pkey; Type: CONSTRAINT; Schema: recitation; Owner: david; Tablespace: 
--

ALTER TABLE ONLY solved
    ADD CONSTRAINT solved_pkey PRIMARY KEY (sid, cid, recitation, problem, letter);


--
-- Name: student_pkey; Type: CONSTRAINT; Schema: recitation; Owner: david; Tablespace: 
--

ALTER TABLE ONLY student
    ADD CONSTRAINT student_pkey PRIMARY KEY (id);


--
-- Name: subproblem_pkey; Type: CONSTRAINT; Schema: recitation; Owner: david; Tablespace: 
--

ALTER TABLE ONLY subproblem
    ADD CONSTRAINT subproblem_pkey PRIMARY KEY (cid, recitation, problem, letter);


--
-- Name: takes_pkey; Type: CONSTRAINT; Schema: recitation; Owner: david; Tablespace: 
--

ALTER TABLE ONLY takes
    ADD CONSTRAINT takes_pkey PRIMARY KEY (cid, sid);


--
-- Name: track_pkey; Type: CONSTRAINT; Schema: recitation; Owner: david; Tablespace: 
--

ALTER TABLE ONLY track
    ADD CONSTRAINT track_pkey PRIMARY KEY (cid, sid, name);


--
-- Name: problem_cid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY problem
    ADD CONSTRAINT problem_cid_fkey FOREIGN KEY (cid, recitation) REFERENCES recitation(cid, name);


--
-- Name: recitation_cid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY recitation
    ADD CONSTRAINT recitation_cid_fkey FOREIGN KEY (cid) REFERENCES course(id);


--
-- Name: solved_cid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY solved
    ADD CONSTRAINT solved_cid_fkey FOREIGN KEY (cid, recitation, problem, letter) REFERENCES subproblem(cid, recitation, problem, letter);


--
-- Name: solved_sid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY solved
    ADD CONSTRAINT solved_sid_fkey FOREIGN KEY (sid) REFERENCES student(id);


--
-- Name: subproblem_cid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY subproblem
    ADD CONSTRAINT subproblem_cid_fkey FOREIGN KEY (cid, recitation, problem) REFERENCES problem(cid, recitation, problem);


--
-- Name: takes_cid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY takes
    ADD CONSTRAINT takes_cid_fkey FOREIGN KEY (cid) REFERENCES course(id);


--
-- Name: takes_sid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY takes
    ADD CONSTRAINT takes_sid_fkey FOREIGN KEY (sid) REFERENCES student(id);


--
-- Name: track_cid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY track
    ADD CONSTRAINT track_cid_fkey FOREIGN KEY (cid, name) REFERENCES recitation(cid, name);


--
-- Name: track_sid_fkey; Type: FK CONSTRAINT; Schema: recitation; Owner: david
--

ALTER TABLE ONLY track
    ADD CONSTRAINT track_sid_fkey FOREIGN KEY (sid) REFERENCES student(id);


--
-- Name: public; Type: ACL; Schema: -; Owner: david
--

REVOKE ALL ON SCHEMA public FROM PUBLIC;
REVOKE ALL ON SCHEMA public FROM david;
GRANT ALL ON SCHEMA public TO david;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--


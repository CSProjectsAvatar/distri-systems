from concurrent.futures import thread
from distutils import filelist
import enum
from fileinput import filelineno
from logging import PlaceHolder
from pstats import Stats
from random import randint
from time import sleep
from turtle import onclick
import streamlit as st
import io
import pandas as pd
import client

import grpc
import middleware_pb2 as mid
import middleware_pb2_grpc as mid_grpc

st.title("Tournament Manager - Stats")

if 'node' not in st.session_state:
    st.session_state.node = client.grpcNode()
    st.session_state.tourNames = []
    st.session_state.tourIds = {}
    # st.session_state.t_ids = []

if 'prevMatch' not in st.session_state:
    st.session_state.prevMatch = 0

node = st.session_state.node
tourNames = st.session_state.tourNames
prevMatchAmm = st.session_state.prevMatch

# tourn_ids = st.session_state.t_ids
tourn_ids = st.session_state.tourIds
# st.write(st.session_state.prevMatch)

prevMatchesDict = {}
# prevMatchs = 0
while True:

    placeholder = st.empty()
    with placeholder.container():
        node.get_all_stats()

        for id, stat in node.tourStats.items():
            statis: mid.StatsResp = stat
            # st.write(statis)
            if id not in prevMatchesDict:
                prevMatchesDict[id] = 0

            prevMatchs = prevMatchesDict[id]
            newMatch = statis.matches

            cont = st.empty()
            with cont.container():
                col1, col2, col3, col4 = st.columns(4)
                col1.metric("TourName", str(statis.tourName))
                # col2.metric("Matches", str(newMatch), str(newMatch - st.session_state.prevMatch))
                col2.metric("Matches", str(newMatch), str(newMatch - prevMatchs))
                col3.metric("Best Player", str(statis.bestPlayer))
                col4.metric("Winner", str(statis.winner))

            prevMatchesDict[id] = newMatch
        sleep(5)
        placeholder.empty()

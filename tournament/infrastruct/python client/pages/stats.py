from concurrent.futures import thread
from distutils import filelist
import enum
from fileinput import filelineno
from logging import PlaceHolder
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

node = st.session_state.node
tourNames = st.session_state.tourNames
# tourn_ids = st.session_state.t_ids
tourn_ids = st.session_state.tourIds

while True:

    placeholder = st.empty()
    with placeholder.container():
        node.get_all_stats()
        # st.write(node.tourStats)

        # for _, stat in enumerate(node.tourStats):
        for id, stat in node.tourStats.items():
            st.write(stat)
            st.write()

            for_run = randint(0,10)
            running = True
            already_run = randint(for_run, for_run+10)
            cont = st.empty()
            with cont.container():
                col1, col2, col3 = st.columns(3)
                col1.metric("For Run", str(for_run) + 'Pcs', "1.2")
                col2.metric("Running", str(running) + 'Pcs', "-8%")
                col3.metric("Already Run", str(already_run) + 'Pcs', "4%")
        sleep(5)
        placeholder.empty()

# get the key, values of a dict
# dict = {'a': 1, 'b': 2, 'c': 3}
# for key, value in dict.items():
#     st.write(key, value)
#     st.write()
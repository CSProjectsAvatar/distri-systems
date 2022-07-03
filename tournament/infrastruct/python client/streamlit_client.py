from concurrent.futures import thread
from distutils import filelist
import enum
from fileinput import filelineno
from logging import PlaceHolder
from time import sleep
from turtle import onclick
import streamlit as st
import io
import pandas as pd
import client

import grpc
import middleware_pb2 as mid
import middleware_pb2_grpc as mid_grpc

# enum tournament type
tour_type = {
    'First Defeat': 0,
    'All vs All': 1,
}

st.title("Tournament Manager")

if 'node' not in st.session_state:
    st.session_state.node = client.grpcNode()
    st.session_state.tourNames = []

node = st.session_state.node
tourNames = st.session_state.tourNames

# Input 'Tournament Name' with button to create a new tournament
new_t_name = st.text_input("Tournament Name", placeholder='Chez Tournament')

if new_t_name != '' and new_t_name not in tourNames:
    tourNames.append(new_t_name)

st.write(tourNames)

# Tournaments Expanders
for t_name in tourNames:
    # Create a new tournament
    with st.expander(t_name):
        files_placeholder = st.empty()
        with files_placeholder.container():
            file_list = []
            # input for get the game
            game = st.file_uploader('Bring the game', ['py'], key='game_' + t_name)
            if game is not None:
                # To read file as bytes:
                bytes_data = game.getvalue()
                # grpc game File
                game_file = mid.File(name='game', data=bytes_data)

                # MOCK
                game_file = 'game.py'
                st.write(game_file)
                file_list.append(game_file)
                st.write(file_list)

            # input for get the players
            players_files = st.file_uploader("Bring the Players", ['py'], accept_multiple_files=True,
                                             key='players_' + t_name)

            for i, player in enumerate(players_files):
                # To read file as bytes:
                bytes_data = player.getvalue()
                # grpc player File
                player_file = mid.File(name='player' + str(i), data=bytes_data)
                file_list.append(player_file)

            selected = st.selectbox('Type of tournament', tour_type, key='tour_type_' + t_name)
            st.write(tour_type[selected])

            if st.button("Upload", key='upBttn_' + t_name):
                # request for the game and players
                # if game is None or len(players_files) < 2:
                #     st.error("Please upload the game and players")
                # else:
                # MOCK TOURNAMENT
                t_name = 'Chez Tournament'
                selected = "First Defeat"
                file_list = [
                    mid.File(name='game', data=b'game.py'),
                    mid.File(name='player0', data=b'player0.py'),
                    mid.File(name='player1', data=b'player1.py'),
                ]

                node.upload_tournment(t_name, tour_type[selected], file_list)

            run_bttn = st.button("Run", key='runBttn_' + t_name)
            st.write(tour_type[selected])
            st.write(file_list)

        if run_bttn:
            files_placeholder.empty()
            st.write("Running the tournament")

            for_run = 20
            running = 5
            already_run = 2

            while True:
                placeholder = st.empty()
                with placeholder.container():
                    col1, col2, col3 = st.columns(3)
                    col1.metric("For Run", str(for_run) + 'Pcs', "1.2 °F")
                    col2.metric("Running", str(running) + 'Pcs', "-8%")
                    col3.metric("Already Run", str(already_run) + 'Pcs', "4%")

                    # sleep 2 seconds
                    for_run += 1
                    running += 1
                    already_run += 2
                    sleep(3)

                placeholder.empty()

# # seleccionar el tipo de torneo

# # resultados y estadisticas

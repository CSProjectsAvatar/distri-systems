from concurrent.futures import thread
from distutils import filelist
import enum
from fileinput import filelineno
from logging import PlaceHolder
from random import randint
from time import sleep
from turtle import onclick
from scipy import rand
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
    'For Groups': 2
}

st.title("Tournament Manager")

if 'node' not in st.session_state:
    st.session_state.node = client.grpcNode()
    st.session_state.tourNames = []
    st.session_state.tourIds = {}
    st.session_state.t_ids = []

node = st.session_state.node
tourNames = st.session_state.tourNames
tour_ids = st.session_state.tourIds
# node = client.grpcNode()
# tourNames = []
# tour_ids = {}
# tourn_ids = st.session_state.t_ids

# Input 'Tournament Name' with button to create a new tournament
new_t_name = st.text_input("Tournament Name", placeholder='Chez Tournament')

if new_t_name != '' and new_t_name not in tourNames:
    tour_ids[new_t_name] = ''
    tourNames.append(new_t_name)

# st.write(tourNames)
# Tournaments Expanders
for t_name in tourNames:
    t_id = tour_ids[t_name]
    st.write(t_id)
    # Create a new tournament
    if t_id != '':
        st.write(t_name + " Uploaded")

    else:
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
                    file_list.append(game_file)

                # input for get the players
                players_files = st.file_uploader("Bring the Players", ['py'], accept_multiple_files=True,
                                                key='players_' + t_name)

                for i, player in enumerate(players_files):
                    # To read file as bytes:
                    bytes_data = player.getvalue()
                    # grpc player File
                    player_file = mid.File(name=f'{player.name}', data=bytes_data)
                    file_list.append(player_file)

                selected = st.selectbox('Type of tournament', tour_type, index=1, key='tour_type_' + t_name)
                # st.write(tour_type[selected])

                # t_id = tourn_ids[i] if i < len(tourn_ids) else None

                if st.button("Upload & Run", key='upBttn_' + t_name):
                    # request for the game and players
                    # if game is None or len(players_files) < 2:
                    #     st.error("Please upload the game and players")
                    # else:
                    # MOCK TOURNAMENT
                    # selected = "First Defeat"
                    # selected = "All vs All"
                    poss_file_lst = [
                        mid.File(name= t_name, data=b'import random as r\nprint(r.randint(1, 3))', is_game=True),
                        mid.File(name='player1', data=b'player1.py'),
                        mid.File(name='player2', data=b'player1.py'),
                        mid.File(name='player3', data=b'player1.py'),
                        mid.File(name='player4', data=b'player1.py'),
                        mid.File(name='player5', data=b'player1.py'),
                        mid.File(name='player6', data=b'player1.py'),
                        mid.File(name='player7', data=b'player1.py'),
                        mid.File(name='player8', data=b'player1.py'),
                        mid.File(name='player9', data=b'player1.py'),
                        mid.File(name='player10', data=b'player1.py'),
                    ]
                    player_number = randint(4, len(poss_file_lst)-2)
                    file_list = poss_file_lst[:player_number+1]

                    resp = node.upload_tournment(t_name, tour_type[selected], file_list)
                    files_placeholder.empty()


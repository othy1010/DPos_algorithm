# DPOS Algorithm
To run the dpos algorithm simulation, first clone the repository and navigate to the root directory of the project. Then, run the following command:

Copy code
```
go run main.go
```
This will execute the main.go file, which contains the simulation for the dpos algorithm.

# Overview
This project is a simulation of the dpos (delegated proof of stake) algorithm. In a dpos system, token holders can delegate their stake to "delegates" who are responsible for validating transactions and maintaining the network. The delegates are chosen in a round-robin fashion, and the weight of each delegate's vote is proportional to the total stake delegated to them.

# Dependencies
This project requires the following dependencies to be installed:

- Go version 1.15 or higher

# How to Use
The simulation can be run by executing the main.go file. This will simulate the dpos algorithm using the default parameters, including the number of delegates and the total number of token holders. These parameters can be modified by modifying the constants at the top of the main.go file.

Once the simulation has been run, the output will include the final state of the network, including the delegates and their respective stake.

# Contributions
Contributions to this project are welcome. If you have an idea for a new feature or have found a bug, please open an issue on the [Github](https://github.com/othy1010/DPos_algorithm/) repository for this project.

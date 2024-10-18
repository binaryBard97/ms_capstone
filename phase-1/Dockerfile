# Use an official Ubuntu as the base image
FROM ubuntu:22.04

# Set environment variables
ENV DEBIAN_FRONTEND=noninteractive

# Install dependencies and wget
RUN apt-get update && apt-get install -y \
    wget \
    apt-transport-https \
    ca-certificates \
    gnupg2 \
    lsb-release \
    && rm -rf /var/lib/apt/lists/*

# Add OpenSUSE repository for CRIU
RUN sh -c "echo 'deb http://download.opensuse.org/repositories/devel:/tools:/criu/xUbuntu_$(lsb_release -rs)/ /' > /etc/apt/sources.list.d/criu.list"

# Add OpenSUSE signing key
RUN wget -qO - https://download.opensuse.org/repositories/devel:/tools:/criu/xUbuntu_$(lsb_release -rs)/Release.key | apt-key add -

# Update package list and install CRIU
RUN apt-get update && apt-get install -y criu

# Verify CRIU installation
RUN criu --version

# Set the default command (optional)
CMD ["bash"]

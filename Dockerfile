FROM fedora:32

# docker login docker.pkg.github.com -u suvam0451
# docker build -t docker.pkg.github.com/suvam0451/video-encoding-kit/suite:latest .
# docker push docker.pkg.github.com/suvam0451/video-encoding-kit/suite:latest

# Proxy setup
RUN sudo echo "proxy=http://172.16.2.30:8080" >> /etc/dnf/dnf.conf


# Packages
RUN sudo dnf install -y https://download1.rpmfusion.org/free/fedora/rpmfusion-free-release-$(rpm -E %fedora).noarch.rpm https://download1.rpmfusion.org/nonfree/fedora/rpmfusion-nonfree-release-$(rpm -E %fedora).noarch.rpm
RUN sudo dnf install -y libXcomposite libXcursor libXi libXtst libXrandr alsa-lib mesa-libEGL libXdamage mesa-libGL libXScrnSaver
RUN sudo dnf install -y ffmpeg youtube-dl aria2 wget unzip

# Package specific proxy
RUN sudo echo "https_proxy=http://172.16.2.30:8080" >> /etc/wgetrc


# Download "drivekit" binary to bin
RUN sudo wget -P /bin https://github.com/suvam0451/video-encoding-kit/releases/latest/download/drivekit && sudo chmod +x /bin/drivekit

# Undo proxy related modifications
RUN sudo sed -i '$d' /etc/dnf/dnf.conf
RUN sudo sed -i '$d' /etc/wgetrc
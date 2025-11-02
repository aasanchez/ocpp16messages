FROM mcr.microsoft.com/devcontainers/base:ubuntu AS builder

ARG USERNAME=vscode
ARG USER_HOME=/home/${USERNAME}
ARG ZSH_DIR=${USER_HOME}/.oh-my-zsh
ENV DEBIAN_FRONTEND=noninteractive
ENV SHELL=/usr/bin/zsh
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

# Install dependencies
RUN apt-get update && apt-get install -y \
    zsh git curl ca-certificates locales tzdata gnupg git-extras catimg ripgrep build-essential \
    zsh locales ca-certificates gnupg git-extras

# Download lazygit
RUN LAZYGIT_VERSION=$(curl -s "https://api.github.com/repos/jesseduffield/lazygit/releases/latest" | grep -Po '"tag_name": *"v\K[^"]*') && \
    curl -Lo lazygit.tar.gz "https://github.com/jesseduffield/lazygit/releases/download/v${LAZYGIT_VERSION}/lazygit_${LAZYGIT_VERSION}_Linux_x86_64.tar.gz" && \
    tar xf lazygit.tar.gz lazygit && \
    install lazygit -D -t /usr/local/bin/ && \
    rm lazygit.tar.gz

# Download other assets and binaries
RUN curl -fsSL https://opencode.ai/install | bash && \
    ln -sf /usr/local/bin/opencode /usr/bin/opencode

# Prepare assets and configs
RUN mkdir -p /app
WORKDIR /app

# Clone Powerlevel10k theme
RUN su - ${USERNAME} -c 'git clone --depth=1 https://github.com/romkatv/powerlevel10k ${ZSH_DIR}/custom/themes/powerlevel10k'

# Copy .zshrc from local directory
COPY .devcontainer/.zshrc /home/${USERNAME}/.zshrc

# Write .zshrc for the user
RUN chown ${USERNAME}:${USERNAME} /home/${USERNAME}/.zshrc

# Copy binaries and assets to a dedicated directory
RUN mkdir -p /build
RUN cp /usr/local/bin/lazygit /build/
RUN cp /usr/local/bin/opencode /build/
RUN cp -r ${ZSH_DIR} /build/.oh-my-zsh

# Final stage
FROM mcr.microsoft.com/devcontainers/base:ubuntu

# Copy only the necessary files from the build stage
COPY --from=builder /build /app

# Set environment variables
ENV ZSH=/app/.oh-my-zsh
ENV SHELL=/usr/bin/zsh
ENV LANG=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8

# Copy the pre-configured .zshrc
COPY .devcontainer/.zshrc /home/vscode/.zshrc

# Set working directory
WORKDIR /app

# Set entrypoint or default command
CMD ["/app/lazygit"]

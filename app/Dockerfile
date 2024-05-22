FROM oven/bun:latest

WORKDIR /usr/src/app

RUN apt-get update && apt-get install -y curl ca-certificates git

RUN curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.7/install.sh | bash \
    && export NVM_DIR="$HOME/.nvm" \
    && [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" \
    && nvm install 22.2.0 \
    && nvm use 22.2.0 \
    && nvm alias default 22.2.0

COPY package.json bun.lockb ./

RUN bun install --frozen-lockfile

COPY . .

CMD ["bash", "-c", "source $HOME/.nvm/nvm.sh && bun --bun run dev"]

FROM ruby:2.7

ENV LANG=C.UTF-8 \
    TZ=Asia/Tokyo

WORKDIR /app
RUN apt-get update -qq && apt-get install -y postgresql-client
# install nodejs(LTS)
RUN curl -fsSL https://deb.nodesource.com/setup_lts.x | bash - && apt-get install -y nodejs
# install yarn
RUN npm install --global yarn

COPY Gemfile /app/Gemfile
COPY Gemfile.lock /app/Gemfile.lock
RUN bundle install

COPY entrypoint.sh /usr/bin/
RUN chmod +x /usr/bin/entrypoint.sh
ENTRYPOINT [ "entrypoint.sh" ]
EXPOSE 3000

CMD ["rails", "server", "-b", "0.0.0.0"]
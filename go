#!/bin/sh

echo "Installing clean project"

mv .git/modules .             && \
rm -rf .git/                  && \
git init                      && \
mv modules .git/              && \
git add styles/brackets.css   && \
git add .gitmodules           && \
rm go

echo "brackets.css project successfully installed"

git status

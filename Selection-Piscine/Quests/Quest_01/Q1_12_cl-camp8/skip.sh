#!//binbash/

# ls | awk 'NR % 2 == 1'

# ls -l | sed '1d'  -n '1~2p'

ls -l | sed '1d' | sed -n 'p;n'
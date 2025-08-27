# Start Python API in background
cd /opt/render/project/src # Render's project path
python -m pip install -r requirements.txt
python main.py &

# Wait a moment for Python API to start
sleep 5

# Start Go application
./app
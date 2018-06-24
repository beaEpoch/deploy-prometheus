# deploy-prometheus

prometheus 系列启动文件 service

systemd 管理服务配置文件

用法：

$ systemctl --user enable /path/to/absolute/xx.service

Created symlink from /home/dev/.config/systemd/user/default.target.wants/xx.service to /home/dev/.config/systemd/user/xx.service.

$ systemctl --user start xx.service

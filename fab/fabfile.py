#!/usr/bin/env python
from invoke import run
from invoke import task
from fabric import Connection
import sys, getopt,time,os,yaml

@task
def deploy(ctx,env):

    run("echo '[Current path] ' `pwd`;")

    data = parse_config("prod")
    app_name = data["app_name"]
    hosts = data["hosts"][env]
    exec_files = data["exec_files"]
    pre_cmd = data["pre_cmd"]
    exec_cmd = data["exec_cmd"]
    port = data["deploy_env"][env]["port"]
    macro = data["macro"]

    # 执行项目本地预编译命令
    for k in pre_cmd:
        print(k)
        run(k)

    t = time.strftime("%Y%m%d%H%M%S", time.localtime())
    back_dir = "{}.{}".format(app_name,t)

    conf_data = open("conf.yml")

    for h in hosts:
        print(h["host"])
        print(h["remote_path"])
        try:

            c = Connection(h["host"],connect_kwargs={"password": h["password"]})
            c.run(f'mkdir -p {h["remote_path"]}')

            file_data = conf_data.read()

            with c.cd(h["remote_path"]):

                c.run("echo '[current path] ' `pwd`; ")
                back_log_path = os.path.join(h["back_log"],app_name,back_dir)

                result1 = c.run(f'ls {h["remote_path"]} | wc -l')
                if int(result1.stdout) > 0:
                    c.run(f'mkdir -p {back_log_path}')
                    c.run(f'mv {h["remote_path"]}/* {back_log_path}')
                    print("back")

                # send_file
                for k in exec_files:
                    arr = k.split(".")
                    if arr[len(arr) - 1] == "yml":
                        result =  c.run("ifconfig eth0|grep inet|grep -v 127.0.0.1|grep -v inet6|awk '{print $2}'|tr -d \"addr:\"")
                        file_data = replace_macro(file_data,result.stdout,t,app_name,str(port),macro)
                        tmp = './conf_ini.yml'
                        f = open(tmp,"a")
                        f.write(file_data)
                        f.close()
                        c.put(tmp,'{}/{}'.format(h["remote_path"],k))
                        os.remove(tmp)
                    else:
                        c.put(k,h["remote_path"])

                # 执行后面命令
                for k in exec_cmd:
                    c.run(k)
        except Exception as e:

            print(e)

            print(result1.stdout)
            if int(result1.stdout) > 0:
                c.run(f'mv {back_log_path}/* {h["remote_path"]}')

    conf_data.close()


@task
def rollback(ctx,env):

    data = parse_config()
    hosts = data["hosts"][env]
    app_name = data["app_name"]

    for h in hosts:
        c = Connection(h["host"],connect_kwargs={"password": h["password"]})
        with c.cd(os.path.join(h["back_log"],app_name)):
            result = c.run("ls |tail -n 1")
            back_file = result.stdout.replace("\n","")
            c.run(f'cp -r {back_file}/* /tmp')



def parse_config():
    load=yaml.load_all(open("fab_config.yml"),Loader=yaml.FullLoader)

    for data in load:
        return data


def replace_macro(file_data,intranet,time_str,app_name,port,macro):

    file_data = file_data.replace("%app_name%",app_name)
    file_data = file_data.replace("%build_date%",time_str)
    file_data = file_data.replace("%intranet%",intranet)
    file_data = file_data.replace("%port%",port)

    for k in macro:
        file_data = file_data.replace('%{}%'.format(k),macro[k])

    return file_data
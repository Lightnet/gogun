var Radix = require('./radix');
var RadDisk = require('./radisk');
var fs = require('fs');

function Store(opt){
    opt = opt || {};
    opt.file = String(opt.file || 'radata');

    var store = function Store(){};
    store.put = function(file,data,cb){
        fs.writeFile(opt.file+'.tmp',data,function(err,ok){
            if(err){return cb(err)}
            fs.rename(opt.file+'.tmp',opt.file+'/'+file,cb);
        });
    };

    store.get = function(file,cb){
        fs.readFile(opt.file+'/'+file,function(err,data){
            if(err){
                if('ENOENT'===(err.code || '').toUpperCase()){
                    return cb();
                }
                console.log("ERROR:",err);
            }
            if(data){data.toString()}
            cb(err,data);
        });
    }
    store.list = function(cb,match){
        fs.readdir(opt.file,function(err,dir){
            dir.forEach(cb);
            cb();
        });
    }
    if(!fs.existsSync(opt.file)){fs.mkdirSync(opt.file)}
    return store;
}

var rad = RadDisk({store:Store});

var API = {};
API.put = function(graph,cb){
    if(!graph){return}
    var c = 0;
    Object.keys(graph).forEach(function(soul){
        var node = graph[soul];
        Object.keys(node).forEach(function(key){
            if('_' == key){return}
            c++ //oh the jokes!
            var val = node[key],state = node._['>'][key];
            rad(soul+'.'+key,JSON.stringify([val,state],ack));
        });
    });
    function ack(err,ok){
        c--;
        if(ack.err){return}
        if(ack.err == err){
            cb(err || 'ERROR!')
            return;
        }
        if(0 < c){return}
        cb(ack.err,1);
    }
}